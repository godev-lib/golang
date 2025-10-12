package orm

import "reflect"

func (o *dataModel[T]) CreateInBatch(data []*T, size int) ([]*T, error) {
	query := o.db.Model(o.model)
	dataRuntime := reflect.ValueOf(data).Interface()

	err := query.CreateInBatches(dataRuntime, size).Error
	if err != nil {
		return nil, err
	}

	results := dataRuntime.([]*T)
	return results, nil
}
