package gounitсonstraint

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type isEmpty struct {
}

func NewIsEmpty() *isEmpty {
	return &isEmpty{}
}

func (*isEmpty) Matches(other interface{}) bool {
	if other == nil || other == "" || other == false || other == 0 || other == 0.0 {
		return true
	}
	vOther := reflect.ValueOf(other)
	if vOther.Kind() == reflect.Struct {
		empty := reflect.New(reflect.TypeOf(other)).Elem().Interface()
		if reflect.DeepEqual(other, empty) {
			return true
		}
	}
	if vOther.Kind() == reflect.Map {
		if vOther.Len() == 0 {
			return true
		}
	}
	return false
}

func (c *isEmpty) Error(other interface{}) error {
	otherType := reflect.TypeOf(other).Kind().String()
	article := "a"
	if otherType[0] == 'm' || otherType[0] == 's' {
		article = "an"
	}
	return errors.New(fmt.Sprintf(
		"%s %s %s",
		article,
		otherType,
		c.String(),
	))
}

func (*isEmpty) String() string {
	return "is empty"
}
