package orm

import (
	"reflect"
	"strings"
)

func (o *dataModel[T]) Delete(filter Filter) error {
	query := o.db.Model(o.model)
	dataRuntime := reflect.New(o.dataType).Interface()

	listConds := []string{}
	args := []interface{}{}
	for _, cond := range filter.Conditions {
		listConds = append(listConds, cond.Query)
		args = append(args, cond.Arg)
	}
	whereStr := strings.Join(listConds, string(filter.OperatorCondition))
	query = query.Where(whereStr, args...)

	if filter.Unscoped {
		query = query.Unscoped()
	}

	err := query.Delete(dataRuntime).Error
	if err != nil {
		return err
	}

	return nil
}
