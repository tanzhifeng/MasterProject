package mysql

import (
	"database/sql"
	"fmt"

	"../../common"
	_ "github.com/go-sql-driver/mysql"
)

type ClientMySQL struct {
	driver        string
	ip            string
	port          uint
	account       string
	password      string
	dbname        string
	maxOpenConn   int
	maxIdleConn   int
	db            *sql.DB
	statements    map[uint]*OperateStmt
	channelWrites map[uint64]chan *Operate
	channelRead   chan *OperateResult
}

func (o *ClientMySQL) handlerOperateSingle(stmt *sql.Stmt, builder QueryBuilder, operate *OperateSingle) *OperateSingleResult {
	rows, err := stmt.Query(operate.OperateArgs...)
	if err != nil {
		fmt.Println(err.Error())
		return &OperateSingleResult{Success: false, OperateArgs: operate.OperateArgs}
	}
	defer rows.Close()

	var res []interface{}
	for rows.Next() {
		fields := builder()
		err = rows.Scan(*fields...)
		if err != nil {
			fmt.Println(err.Error())
			return &OperateSingleResult{Success: false, QueryRows: res, OperateArgs: operate.OperateArgs}
		}

		res = append(res, fields)
	}

	return &OperateSingleResult{Success: true, QueryRows: res, OperateArgs: operate.OperateArgs}
}

func (o *ClientMySQL) handlerClientWrite(userid uint64) {
	channelWrite := o.channelWrites[userid]

	for {
		v, ok := <-channelWrite
		if !ok {
			return
		}
		o.db.Ping()

		result := &OperateResult{OperateType: v.OperateType, Success: true, OperateOptions: v.OperateOptions, Handler: v.Handler}

		switch v.OperateType {
		case OperateCommand:
			for i := 0; i < len(v.OperateList); i++ {
				operate := v.OperateList[i]
				opStmt := o.statements[operate.OperateID]
				resultSingle := o.handlerOperateSingle(opStmt.Stmt, opStmt.Builder, operate)

				result.Success = result.Success && resultSingle.Success
				result.Results = append(result.Results, resultSingle)
			}
		case OperateTransaction:
			tx, err := o.db.Begin()
			if err != nil {
				for i := 0; i < len(v.OperateList); i++ {
					operate := v.OperateList[i]
					result.Results = append(result.Results, &OperateSingleResult{Success: false, OperateArgs: operate.OperateArgs})
				}
				result.Success = false
			} else {
				for i := 0; i < len(v.OperateList); i++ {
					operate := v.OperateList[i]
					opStmt := o.statements[operate.OperateID]
					resultSingle := o.handlerOperateSingle(tx.Stmt(opStmt.Stmt), opStmt.Builder, operate)

					result.Success = result.Success && resultSingle.Success
					result.Results = append(result.Results, resultSingle)
				}

				if result.Success {
					tx.Commit()
				} else {
					tx.Rollback()
				}
			}
		default:
			result.Success = false
			fmt.Println("Unmatch OperateType :", v.OperateType)
		}
		o.channelRead <- result
	}
}

func (o *ClientMySQL) Start(driver string, ip string, port uint, account string, password string, dbname string, sqls map[uint]*OperateSQL, maxOpenConn int, maxIdleConn int) (chan *OperateResult, error) {
	o.driver = driver
	o.ip = ip
	o.port = port
	o.account = account
	o.password = password
	o.dbname = dbname
	o.maxOpenConn = maxOpenConn
	o.maxIdleConn = maxIdleConn
	o.statements = make(map[uint]*OperateStmt)
	o.channelRead = make(chan *OperateResult, common.DefaultChannelLen)
	o.channelWrites = make(map[uint64]chan *Operate)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", o.account, o.password, o.ip, o.port, o.dbname)
	db, err := sql.Open(o.driver, dataSource)
	if err != nil {
		fmt.Println(err.Error())

		o.clean()
		return nil, err
	}

	o.db = db
	o.db.SetMaxOpenConns(o.maxOpenConn)
	o.db.SetMaxIdleConns(o.maxIdleConn)

	for k, v := range sqls {
		stmt, err := o.db.Prepare(v.SQL)
		if err != nil {
			fmt.Println(err.Error())

			o.db.Close()
			o.clean()
			return nil, err
		}
		o.statements[k] = &OperateStmt{Stmt: stmt, Builder: v.Builder}
	}

	return o.channelRead, err
}

func (o *ClientMySQL) Stop() {
	o.db.Close()

	o.clean()
}

func (o *ClientMySQL) clean() {
	for _, v := range o.statements {
		v.Stmt.Close()
	}

	for _, v := range o.channelWrites {
		close(v)
	}

	o.db = nil
	close(o.channelRead)
	o.statements = nil
	o.channelWrites = nil
}

func (o *ClientMySQL) CreateClientChannel(userid uint64) {
	_, ok := o.channelWrites[userid]
	if !ok {
		o.channelWrites[userid] = make(chan *Operate, common.DefaultChannelLen)
		go o.handlerClientWrite(userid)
	}
}

func (o *ClientMySQL) DeleteClientChannel(userid uint64) {
	close(o.channelWrites[userid])
	delete(o.channelWrites, userid)
}

func (o *ClientMySQL) Execute(userid uint64, operate *Operate) {
	channelWrite, ok := o.channelWrites[userid]

	if ok {
		channelWrite <- operate
	}
}
