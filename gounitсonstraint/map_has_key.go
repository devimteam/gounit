package gounitсonstraint

import (
	"fmt"
	"reflect"
    "github.com/pkg/errors"
)

type mapHasKey struct {
	key interface{}
}

func NewMapHasKey(key interface{}) *mapHasKey {
	return &mapHasKey{key: key}
}

func (c *mapHasKey) Matches(other interface{}) bool {
	vOther := reflect.ValueOf(other)
	vKey := reflect.ValueOf(c.key)
	vIndex := vOther.MapIndex(vKey)
	if vIndex.IsValid() {
		return true
	}
	return false
}

func (c *mapHasKey) Error(other interface{}) error {
    return errors.New("an map " + c.String())
}

func (c *mapHasKey) String() string {
	return fmt.Sprintf("has the key %#v", c.key)
}
