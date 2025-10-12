package orm

import "reflect"

func (o *dataModel[T]) CreateOne(data *T) (*T, error) {
	query := o.db.Model(o.model)
	dataRuntime := reflect.ValueOf(data)
	err := query.Create(dataRuntime.Interface()).Error

	if err != nil {
		return nil, err
	}

	result := dataRuntime.Interface().(*T)
	return result, nil
}
