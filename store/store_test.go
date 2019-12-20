package store

import (
	"testing"
)

type Item struct {
	Field1 string
	Field2 int
}

func TestInMemoryStore(t *testing.T) {
	item1 := Item{
		Field1: "one",
		Field2: 1,
	}

	s := NewInMemoryStore()
	s.Store("key", item1)

	var item2 Item
	s.Load("key", &item2)

	if item1.Field1 != item2.Field1 {
		t.Errorf("expect \"%s\" but received \"%s\"", item1.Field1, item2.Field1)
	}

	if item1.Field2 != item2.Field2 {
		t.Errorf("expect %d but received %d", item1.Field2, item2.Field2)
	}

}
