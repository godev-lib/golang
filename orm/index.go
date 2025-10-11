package orm

import (
	"reflect"

	"gorm.io/gorm"
)

type (
	dataModel[T any] struct {
		db               *gorm.DB
		modelRuntime     interface{}
		dataType         reflect.Type
		dataRuntime      reflect.Value
		sliceDataRuntime reflect.Value
	}
	DataMethod[T any] interface {
		Find(filter FindFilter) ([]T, error)
		FindAndCount(filter FindFilter) ([]T, int64, error)
		Count(filter FindFilter) (int64, error)
	}
)

func NewOrm[T any](db *gorm.DB) DataMethod[T] {
	var dataType T
	dataTypeRuntime := reflect.TypeOf(dataType)

	modelRuntime := reflect.New(dataTypeRuntime).Interface()
	dataRuntime := reflect.New(dataTypeRuntime)
	sliceDataRuntime := reflect.New(reflect.SliceOf(dataTypeRuntime))

	return &dataModel[T]{
		db:               db,
		modelRuntime:     modelRuntime,
		dataType:         dataTypeRuntime,
		dataRuntime:      dataRuntime,
		sliceDataRuntime: sliceDataRuntime,
	}
}
