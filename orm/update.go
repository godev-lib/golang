package orm

import (
	"reflect"
)

func (o *dataModel[T]) Update(data *T, filter Filter) (*T, error) {
	dataRuntime := reflect.ValueOf(data).Interface()
	query := o.db.Model(o.model)

	query = query.Where(queryBuilder(filter))

	err := query.Updates(&dataRuntime).Error
	if err != nil {
		return nil, err
	}

	result := dataRuntime.(*T)
	return result, nil
}
