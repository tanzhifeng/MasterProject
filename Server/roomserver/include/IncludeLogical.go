package include

//LogicalHandler ...
type LogicalHandler func(params ...interface{})

type ILogical interface {
	AppendEvent(event int, options ...interface{})
}