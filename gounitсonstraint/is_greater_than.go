package gounitсonstraint

import "fmt"

type isGreaterThan struct {
	expected int
}

func NewIsGreaterThan(expected int) *isGreaterThan {
	return &isGreaterThan{expected: expected}
}

func (c *isGreaterThan) Matches(other interface{}) bool {
	if other.(int) > c.expected {
		return true
	}
	return false
}

func (*isGreaterThan) Error(other interface{}) error {
	return nil
}

func (c *isGreaterThan) String() string {
	return fmt.Sprintf("is greater than %d", c.expected)
}
