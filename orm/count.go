package orm

import "strings"

func (o *dataModel[T]) Count(filter FindFilter) (int64, error) {
	var count int64
	query := o.db.Model(o.modelRuntime)

	listFields := []string{}
	args := []interface{}{}
	for _, item := range filter.Conditions {
		listFields = append(listFields, item.Query)
		args = append(args, item.Arg)
	}

	queryWhere := strings.Join(listFields, string(filter.OperatorCondition))
	query = query.Where(queryWhere, args...)

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
