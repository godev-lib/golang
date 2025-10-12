package orm

import (
	"reflect"
	"strings"
)

func (o *dataModel[T]) Update(data *T, filter Filter) (*T, error) {
	dataRuntime := reflect.ValueOf(data).Interface()
	query := o.db.Model(o.model)

	listConds := []string{}
	args := []interface{}{}
	for _, cond := range filter.Conditions {
		listConds = append(listConds, cond.Query)
		args = append(args, cond.Arg)
	}
	whereStr := strings.Join(listConds, string(filter.OperatorCondition))
	query = query.Where(whereStr, args...)

	err := query.Updates(&dataRuntime).Error
	if err != nil {
		return nil, err
	}

	result := dataRuntime.(*T)
	return result, nil
}
