package orm

import (
	"reflect"
	"strings"
)

func (o *dataModel[T]) Find(filter Filter) ([]T, error) {
	datas := reflect.New(reflect.SliceOf(o.dataType))
	query := o.db.Model(o.model)

	if filter.Unscoped {
		query = query.Unscoped()
	}

	listFields := []string{}
	args := []interface{}{}
	for _, item := range filter.Conditions {
		listFields = append(listFields, item.Query)
		args = append(args, item.Arg)
	}

	queryWhere := strings.Join(listFields, string(filter.OperatorCondition))
	query = query.Where(queryWhere, args...)

	query = query.
		Limit(filter.Limit).
		Offset(filter.Offset)

	err := query.Find(datas.Interface()).Error
	if err != nil {
		return nil, err
	}

	results := datas.Elem().Interface().([]T)
	return results, nil
}
