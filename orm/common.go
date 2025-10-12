package orm

import "strings"

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

func queryBuilder(filter Filter) (string, []interface{}) {
	listFields := []string{}
	args := []interface{}{}

	for _, item := range filter.Conditions {
		listFields = append(listFields, item.Query)
		args = append(args, item.Arg)
	}

	queryWhere := strings.Join(listFields, string(filter.OperatorCondition))
	return queryWhere, args
}
