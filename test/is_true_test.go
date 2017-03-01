package test

import (
	"fmt"
	"testing"

	"github.com/l-vitaly/gounit"
	"github.com/l-vitaly/gounit/gounitсonstraint"
)

type trueTestData struct {
	value  interface{}
	isFail bool
}

var trueTestCase = []trueTestData{
	{true, false},
	{false, true},
	{nil, true},
	{"world", true},
	{123, true},
}

func TestAssertTrue(t *testing.T) {
	u := gounit.New(t)

	trueConstraint := gounitсonstraint.NewIsTrue()

	u.TestCase(trueTestCase, func(i interface{}) {
		tc := i.(trueTestData)
		res := trueConstraint.Matches(tc.value)
		if tc.isFail {
			u.AssertFalse(res, "value:"+fmt.Sprintf("%v", tc.value))
		} else {
			u.AssertTrue(res, "")
		}
	})
}
