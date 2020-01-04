package mysqlcli

import (
	"fmt"
	"sync"

	"../../common/mysql"
)

type MySQLClient struct {
	conn   *mysql.ClientMySQL
	map4SQL map[uint]*mysql.OperateSQL

	channelStop chan bool
	channelRead chan *mysql.OperateResult
}

var mysqlC *MySQLClient
var mysqlCOnce sync.Once

func GetMySQLClient() *MySQLClient {
	mysqlCOnce.Do(func() {
		mysqlC = &MySQLClient{}
		mysqlC.init()
	})
	return mysqlC
}

func (o *MySQLClient) init() {
	o.map4SQL = make(map[uint]*mysql.OperateSQL)

	o.map4SQL[DB_USER_SAVE] = &mysql.OperateSQL{
		SQL: "update bqd.user set gold = ? where userid = ?",
	}

	o.map4SQL[DB_USER_LOGIN] = &mysql.OperateSQL{
		SQL: "select * from bqd.user where account = ? and password = ?",
		Builder: func() *[]interface{} { return &[]interface{}{new(uint64), new(string), new(string), new(uint64)} },
	}

	o.map4SQL[DB_USER_REGISTER] = &mysql.OperateSQL{
		SQL: "insert into bqd.user (account, password) values (?,?)",
	}

	o.map4SQL[DB_USER_SELECT_LAST] = &mysql.OperateSQL{
		SQL:     "select * from bqd.user where userid = @@identity",
		Builder: func() *[]interface{} { return &[]interface{}{new(uint64), new(string), new(string), new(uint64)} },
	}
}

func (o *MySQLClient) Start(ip string, port uint, account string, password string, dbname string, maxOpenConn int, maxIdleConn int) bool {
	o.conn = new(mysql.ClientMySQL)
	channel, err := o.conn.Start("mysql", ip, port, account, password, dbname, o.map4SQL, maxOpenConn, maxIdleConn)

	if err != nil {
		fmt.Println("Start MySQLClient Error :", err.Error())
		return false
	}

	o.channelStop = make(chan bool)
	o.channelRead = channel

	go o.HandleUpdate()

	fmt.Println("MySQLClient Start Successed !")

	return true
}

func (o *MySQLClient) Stop() {
	o.conn.Stop()

	<-o.channelStop
}

func (o *MySQLClient) CreateClientChannel(userid uint64) {
	o.conn.CreateClientChannel(userid)
}

func (o *MySQLClient) DeleteClientChannel(userid uint64) {
	o.conn.DeleteClientChannel(userid)
}

func (o *MySQLClient) Execute(userid uint64, operate *mysql.Operate) {
	o.conn.Execute(userid, operate)
}

func (o *MySQLClient) HandleUpdate() {
	defer close(o.channelStop)

	for {
		v, ok := <-o.channelRead
		if !ok {
			return
		}

		v.Handler(v)
	}
}
