package test

import (
	"fmt"
	"testing"

	"github.com/l-vitaly/gounit"
	"github.com/l-vitaly/gounit/gounitсonstraint"
)

type falseTestData struct {
	value  interface{}
	isFail bool
}

var falseTestCase = []falseTestData{
	{true, true},
	{false, false},
	{nil, true},
	{"world", true},
	{123, true},
}

func TestAssertFalse(t *testing.T) {
	u := gounit.New(t)

	falseConstraint := gounitсonstraint.NewIsFalse()

	u.TestCase(falseTestCase, func(i interface{}) {
		tc := i.(falseTestData)
		res := falseConstraint.Matches(tc.value)
		if tc.isFail {
			u.AssertFalse(res, "value:"+fmt.Sprintf("%v", tc.value))
		} else {
			u.AssertTrue(res, "")
		}
	})
}
