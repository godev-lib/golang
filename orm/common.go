package orm

type OPERATOR string

const (
	AND OPERATOR = "AND"
	OR  OPERATOR = "OR"
)

type Filter struct {
	Limit             int
	Offset            int
	OperatorCondition OPERATOR
	Conditions        []Condition
	Unscoped          bool
}

type Condition struct {
	Query string
	Arg   interface{}
}
