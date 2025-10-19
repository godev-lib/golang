package orm

func (o *dataModel[T]) Count(filter Filter) (int64, error) {
	var count int64
	query := o.db.Model(o.model)

	if filter.Unscoped {
		query = query.Unscoped()
	}

	if len(filter.Conditions) > 0 {
		query = query.Where(queryBuilder(filter))
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
