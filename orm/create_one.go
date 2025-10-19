package orm

import "reflect"

func (o *dataModel[T]) CreateOne(data *T) (*T, error) {
	query := o.db.Model(o.model)
	dataRuntime := reflect.ValueOf(data).Interface()
	err := query.Create(dataRuntime).Error

	if err != nil {
		return nil, err
	}

	result := dataRuntime.(*T)
	return result, nil
}
