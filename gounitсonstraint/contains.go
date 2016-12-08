package gounitсonstraint

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

type contains struct {
	value interface{}
}

func NewContains(value interface{}) *contains {
	return &contains{value: value}
}

func (c *contains) Matches(other interface{}) bool {
	if reflect.TypeOf(other).Kind() != reflect.Slice {
		panic("AssertContains expects is slice")
	}

	vOther := reflect.ValueOf(other)
	vNeedle := reflect.ValueOf(c.value)
	for i := 0; i < vOther.Len(); i++ {
		if reflect.DeepEqual(vOther.Index(i).Interface(), vNeedle.Interface()) {
			return true
		}
	}
	return false
}

func (c *contains) Error(other interface{}) error {
	return errors.New("an slice " + c.String())
}

func (c *contains) String() string {
	return fmt.Sprintf("contains %#v", c.value)
}
