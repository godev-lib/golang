package orm

import (
	"reflect"

	"gorm.io/gorm"
)

type (
	dataModel[T any] struct {
		db       *gorm.DB
		dataType reflect.Type
		model    any
	}
	DataMethod[T any] interface {
		Find(filter Filter) ([]T, error)
		FindOne(filter Filter) (*T, error)
		FindAndCount(filter Filter) ([]T, int64, error)
		Count(filter Filter) (int64, error)
		CreateOne(*T) (*T, error)
		CreateMany([]*T) ([]*T, error)
		CreateInBatch(data []*T, size int) ([]*T, error)
		Update(data *T, filter Filter) (*T, error)
		Delete(filter Filter) error
	}
)

func NewOrm[T any](db *gorm.DB) DataMethod[T] {
	dataType := reflect.TypeFor[T]()
	model := reflect.New(dataType).Interface()

	return &dataModel[T]{
		db:       db,
		dataType: dataType,
		model:    model,
	}
}
