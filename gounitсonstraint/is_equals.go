package gounitсonstraint

import (
    "reflect"
    "fmt"
)

type isEquals struct {
    value interface{}
}

func NewIsEquals(value interface{}) *isEquals {
    return &isEquals{value: value}
}

func (c *isEquals) Matches(other interface{}) bool {
    return reflect.DeepEqual(c.value, other)
}

func (*isEquals) Error(other interface{}) error {
    return nil
}

func (c *isEquals) String() string {
    return fmt.Sprintf("is equal to %#v", c.value);
}
