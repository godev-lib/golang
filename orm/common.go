package orm

type OPERATOR string

const (
	AND OPERATOR = "AND"
	OR  OPERATOR = "OR"
)

type FindFilter struct {
	Limit             int
	Offset            int
	OperatorCondition OPERATOR
	Conditions        []FindCondition
}

type FindCondition struct {
	Query string
	Arg   interface{}
}
