package gounitсonstraint

import (
    "github.com/pkg/errors"
    "fmt"
    "reflect"
)

type count struct {
    expectedCount int
}

func NewCount(expectedCount int) *count {
    return &count{expectedCount: expectedCount}
}

func (c *count) Matches(other interface{}) bool {
    if reflect.TypeOf(other).Kind() != reflect.Slice {
        panic("AssertCount expects is slice")
    }
    vOther := reflect.ValueOf(other)
    if vOther.Len() == c.expectedCount {
        return true
    }
    return false
}

func (c *count) Error(other interface{}) error {
    vOther := reflect.ValueOf(other)
    return errors.New(fmt.Sprintf("actual size %d matches expected size %d", vOther.Len(), c.expectedCount))
}

func (c *count) String() string {
    return fmt.Sprintf("count matches %d", c.expectedCount);
}
