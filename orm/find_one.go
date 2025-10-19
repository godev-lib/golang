package orm

import "reflect"

func (o *dataModel[T]) FindOne(filter Filter) (*T, error) {
	query := o.db.Model(o.model)
	dataRuntime := reflect.New(o.dataType).Interface()

	if len(filter.Conditions) > 0 {
		query = query.Where(queryBuilder(filter))
	}

	err := query.First(&dataRuntime).Error
	if err != nil {
		return nil, err
	}

	result := dataRuntime.(*T)
	return result, nil
}
