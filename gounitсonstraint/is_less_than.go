package gounitсonstraint

import "fmt"

type isLessThan struct {
	expected int
}

func NewIsLessThan(expected int) *isLessThan {
	return &isLessThan{expected: expected}
}

func (c *isLessThan) Matches(other interface{}) bool {
	if other.(int) < c.expected {
		return true
	}
	return false
}

func (*isLessThan) Error(other interface{}) error {
	return nil
}

func (c *isLessThan) String() string {
	return fmt.Sprintf("is less than %d", c.expected)
}
