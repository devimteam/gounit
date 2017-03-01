package test

import (
	"fmt"
	"testing"

	"github.com/l-vitaly/gounit"
	"github.com/l-vitaly/gounit/gounitсonstraint"
)

type nilTestData struct {
	value  interface{}
	isFail bool
}

var nilTestCase = []trueTestData{
	{nil, false},
	{"hello", true},
	{100, true},
	{true, true},
	{1.0, true},
}

func TestAssertNil(t *testing.T) {
	u := gounit.New(t)

	nilConstraint := gounitсonstraint.NewIsNil()

	u.TestCase(nilTestCase, func(i interface{}) {
		tc := i.(nilTestData)
		res := nilConstraint.Matches(tc.value)
		if tc.isFail {
			u.AssertFalse(res, "value:"+fmt.Sprintf("%v", tc.value))
		} else {
			u.AssertTrue(res, "")
		}
	})
}
