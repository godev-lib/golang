package orm

import "reflect"

func (o *dataModel[T]) CreateMany(data []*T) ([]*T, error) {
	query := o.db.Model(o.model)
	dataRuntime := reflect.ValueOf(data).Interface()

	err := query.Create(dataRuntime).Error
	if err != nil {
		return nil, err
	}

	results := dataRuntime.([]*T)
	return results, nil
}
