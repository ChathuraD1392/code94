package inmem

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Repository[T any] interface {
	Add(t *T) error
	Update(id uint, t T) error
	Get(id uint) (T, error)
	GetAll() []T
	Filter(field string, value any) []T
}

var (
	ErrNotFound error
)

func init() {
	ErrNotFound = errors.New("element not found")
}

type InMemoryRepository[T any] struct {
	data  map[uint]T
	mutex sync.RWMutex
	idSeq uint
}

func NewInMemoryRepository[T any]() *InMemoryRepository[T] {
	return &InMemoryRepository[T]{
		data: make(map[uint]T),
	}
}

func (repo *InMemoryRepository[T]) Add(t *T) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.idSeq++
	if err := setField(t, "Id", repo.idSeq); err != nil {
		return err
	}
	if err := setField(t, "CreatedAt", time.Now()); err != nil {
		return err
	}
	if err := setField(t, "UpdatedAt", time.Now()); err != nil {
		return err
	}
	repo.data[repo.idSeq] = *t
	return nil
}

func (repo *InMemoryRepository[T]) Update(id uint, t T) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var (
		old    T
		isExit bool
	)

	old, isExit = repo.data[id]
	if !isExit {
		return ErrNotFound
	}

	createdAt, err := getField(old, "CreatedAt")
	if err != nil {
		return err
	}

	if err := setField(&t, "Id", repo.idSeq); err != nil {
		return err
	}
	if err := setField(&t, "CreatedAt", createdAt); err != nil {
		return err
	}
	if err := setField(&t, "UpdatedAt", time.Now()); err != nil {
		return err
	}
	repo.data[id] = t
	return nil
}

func (repo *InMemoryRepository[T]) Get(id uint) (T, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	ele, isExit := repo.data[id]
	if !isExit {
		return ele, ErrNotFound
	}
	return ele, nil
}

func (repo *InMemoryRepository[T]) GetAll() []T {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	results := make([]T, 0, len(repo.data))
	for _, value := range repo.data {
		results = append(results, value)
	}
	return results
}

func (repo *InMemoryRepository[T]) Filter(field string, value any) []T {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	results := []T{}
	for _, item := range repo.data {
		val := reflect.ValueOf(item)
		fieldValue := val.FieldByName(field)
		if !fieldValue.IsValid() {
			continue
		}

		if reflect.DeepEqual(fieldValue.Interface(), value) {
			results = append(results, item)
		}
	}
	return results
}

func getField(obj interface{}, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("provided object is not a struct")
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("no such field: %s", fieldName)
	}

	return field.Interface(), nil
}

func setField(obj interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("provided object must be a pointer to a struct")
	}

	structValue := v.Elem()
	field := structValue.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s", fieldName)
	}

	fieldValue := reflect.ValueOf(value)
	if field.Type() != fieldValue.Type() {
		return fmt.Errorf("provided value type does not match field type")
	}

	field.Set(fieldValue)
	return nil
}
