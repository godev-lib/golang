package orm

import (
	"context"
	"strings"
)

func (o *dataModel[T]) Delete(filter Filter) error {
	ctx := context.Background()
	query := o.db.Model(o.model)

	listConds := []string{}
	args := []interface{}{}
	for _, cond := range filter.Conditions {
		listConds = append(listConds, cond.Query)
		args = append(args, cond.Arg)
	}
	whereStr := strings.Join(listConds, string(filter.OperatorCondition))
	query = query.Where(whereStr, args...)

	if filter.Unscoped {
		query = query.Unscoped()
	}

	err := query.Delete(ctx).Error
	if err != nil {
		return err
	}

	return nil
}
