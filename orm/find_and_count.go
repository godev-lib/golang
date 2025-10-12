package orm

import "errors"

func (o *dataModel[T]) FindAndCount(filter Filter) ([]T, int64, error) {
	var chanList, chanCount = make(chan []T, 1), make(chan int64, 1)
	var chanErrList, chanErrCount = make(chan error, 1), make(chan error, 1)

	go func() {
		list, err := o.Find(filter)
		chanList <- list
		chanErrList <- err
	}()

	go func() {
		count, err := o.Count(filter)
		chanCount <- count
		chanErrCount <- err
	}()

	if err := <-chanErrCount; err != nil {
		return nil, 0, err
	}

	if err := <-chanErrList; err != nil {
		return nil, 0, err
	}

	count := <-chanCount
	if count == 0 {
		return nil, 0, errors.New("data not found")
	}

	list := <-chanList

	return list, count, nil
}
