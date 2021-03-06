package store

import (
	"errors"
	"reflect"
	"sync"
)

var (
	// ErrNotFound indicate key not exists in storage
	ErrNotFound = errors.New("not found")
)

// Store define simple key-value store function
type Store interface {
	Store(key string, value interface{}) error
	Load(key string, retValue interface{}) error
	LoadOrStore(key string, retValue interface{}) (bool, error)
	Delete(key string) error
}

// default implementation of store
// using in-memory data
type inMemoryStore struct {
	storage sync.Map
}

func (s *inMemoryStore) Store(key string, value interface{}) error {
	s.storage.Store(key, value)

	return nil
}

func (s *inMemoryStore) Load(key string, retValue interface{}) error {
	v, ok := s.storage.Load(key)
	if !ok {
		return ErrNotFound
	}

	if retValue != nil {
		cloneValue(v, retValue)
	}

	return nil
}
func (s *inMemoryStore) LoadOrStore(key string, retValue interface{}) (bool, error) {
	// try to load
	err := s.Load(key, retValue)
	if err != nil {
		// return error
		if err != ErrNotFound {
			return false, err
		}

		// store
		err = s.Store(key, retValue)
		return false, err
	}

	return true, nil
}

func (s *inMemoryStore) Delete(key string) error {
	s.storage.Delete(key)

	return nil
}

func NewInMemoryStore() Store {
	return &inMemoryStore{}
}

func cloneValue(source interface{}, destin interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(destin).Elem().Set(y.Elem())
	} else {
		reflect.ValueOf(destin).Elem().Set(reflect.ValueOf(source))
	}
}
