package orm

import (
	"reflect"
)

func (o *dataModel[T]) Find(filter Filter) ([]T, error) {
	datas := reflect.New(reflect.SliceOf(o.dataType))
	query := o.db.Model(o.model)

	if filter.Unscoped {
		query = query.Unscoped()
	}

	query = query.Where(queryBuilder(filter))

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
