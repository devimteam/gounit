package gounitсonstraint

import "reflect"

type isNil struct {
}

func NewIsNil() *isNil {
	return &isNil{}
}

func (*isNil) Matches(other interface{}) bool {
	obj := reflect.ValueOf(other)
	if !obj.IsValid() {
		return true
	}
	if obj.IsValid() && obj.Kind() != reflect.Ptr {
		return false
	}
	return obj.IsNil()
}

func (*isNil) Error(other interface{}) error {
	return nil
}

func (*isNil) String() string {
	return "is nil"
}
