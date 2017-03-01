package test

import (
	"fmt"
	"testing"

	"github.com/l-vitaly/gounit"
	"github.com/l-vitaly/gounit/gounitсonstraint"
)

type notNilTestData struct {
	value  interface{}
	isFail bool
}

var ptrValue = 1

var notNilTestCase = []notNilTestData{
	{nil, true},
	{"hello", false},
	{100, false},
	{true, false},
	{1.0, false},
    {&ptrValue, false},
}

func TestAssertNotNil(t *testing.T) {
	u := gounit.New(t)

	notNilConstraint := gounitсonstraint.NewNot(gounitсonstraint.NewIsNil())

	u.TestCase(notNilTestCase, func(i interface{}) {
		tc := i.(notNilTestData)
		res := notNilConstraint.Matches(tc.value)
		msg := "value:" + fmt.Sprintf("%v", tc.value)
		if tc.isFail {
			u.AssertFalse(res, msg)
		} else {
			u.AssertTrue(res, msg)
		}
	})
}
