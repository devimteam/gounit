package gounitсonstraint

import "strings"

type not struct {
	constraint Constraint
}

func NewNot(constraint Constraint) *not {
	return &not{constraint: constraint}
}

func (c *not) Matches(other interface{}) bool {
	if !c.constraint.Matches(other) {
		return true
	}
	return false
}

func (*not) negateString(s string) string {
	r := strings.NewReplacer(
        "contains ", "does not contain ",
        "exists ", "does not exists ",
        "has ", "does not have",
        "matches ", "does not match ",
		"is ", "is not ",
		"are ", "are not ",
        "not not ", "not ",
	)
	return r.Replace(s)
}

func (*not) Error(other interface{}) error {
	return nil
}

func (c *not) String() string {
	return c.negateString(c.constraint.String())
}
