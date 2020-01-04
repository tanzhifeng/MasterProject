package mysql

import (
	"database/sql"
)

//operate type
const (
	OperateCommand = iota
	OperateTransaction
)

//QueryBuilder ...
type QueryBuilder func() *[]interface{}

//OperateSQL ...
type OperateSQL struct {
	SQL     string
	Builder QueryBuilder
}

//OperateStmt ...
type OperateStmt struct {
	Stmt    *sql.Stmt
	Builder QueryBuilder
}

//OperateSingle ...
type OperateSingle struct {
	OperateID   uint
	OperateArgs []interface{}
}

//OperateSingleResult ...
type OperateSingleResult struct {
	Success     bool
	QueryRows   []interface{}
	OperateArgs []interface{}
}

//Operate ...
type Operate struct {
	OperateType    uint
	OperateList    []*OperateSingle
	OperateOptions []interface{}
	Handler        HandlerMySQL
}

//OperateResult ...
type OperateResult struct {
	OperateType    uint
	Success        bool
	Results        []*OperateSingleResult
	OperateOptions []interface{}
	Handler        HandlerMySQL
}

type HandlerMySQL func(*OperateResult)