package orm

import "gorm.io/gorm"

func (o *dataModel[T]) GetDB() *gorm.DB {
	return o.db
}
