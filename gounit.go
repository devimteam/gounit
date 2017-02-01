package gounit

import (
	"reflect"
	"testing"

	"github.com/l-vitaly/gounit/gounitсonstraint"
)

type TestCaseFunc func(i interface{})

const (
	// EmptyMessage not show message for assert
	EmptyMessage = ""
)

// GoUnit assert helper
type GoUnit interface {
	// AssertArrayHasKey Reports an error identified by "message" if "array" does
	// not have the "key".
	AssertArrayHasKey(key interface{}, array interface{}, message string)
	// AssertContains Reports an error identified by "message" if "needle" is not
	// an element of "array".
	AssertContains(value interface{}, array interface{}, message string)
	// Reports an error identified by "message" if the number of elements in "array"
	// is not "expectedCount".
	AssertCount(count int, array interface{}, message string)
	// AssertDirectoryExists Reports an error identified by "message" if the
	// directory specified by "dir" does not exist.
	AssertDirectoryExists(dir string, message string)
	// AssertFileExists Reports an error identified by "message" if the file specified
	// by "file" does not exist.
	AssertFileExists(file string, message string)
	// AssertEmpty Reports an error identified by "message" if "actual" is not empty.
	AssertEmpty(actual interface{}, message string)
	// AssertEquals Reports an error identified by "message" if the two variables
	// "expected" and "actual" are not equal.
	AssertEquals(expected interface{}, actual interface{}, message string)
	// AssertTrue Reports an error identified by "message" if "cond" is false.
	AssertTrue(cond bool, message string)
	// AssertFalse Reports an error identified by "message" if "cond" is true.
	AssertFalse(cond bool, message string)
	// AssertFileEquals Reports an error identified by "message: if the file specified
	// by "expected" does not have the same contents as the file specified by "actual".
	AssertFileEquals(expected string, actual string, message string)
	// AssertGreaterThan Reports an error identified by "message" if the value of
	// "actual" is not greater than the value of "expected".
	AssertGreaterThan(expected, actual int, message string)
	// AssertGreaterThanOrEqual Reports an error identified by "message" if the
	// value of "actual" is not greater than or equal to the value of "expected".
	AssertGreaterThanOrEqual(expected, actual int, message string)
	// AssertLessThan Reports an error identified by "message" if the value of
	// "actual" is not less than the value of "expected".
	AssertLessThan(expected, actual int, message string)
	// AssertLessThanOrEqual Reports an error identified by "message" if the value of
	// "actual" is not less than or equal to the value of "expected".
	AssertLessThanOrEqual(expected, actual int, message string)
	//// AssertError Reports an error identified by "message" if "err" is error type
	AssertError(err error, message string)
	// AssertNotError Reports an error identified by "message" if "err" is not error type
	AssertNotError(err error, message string)
	// AssertIsNil Reports an error identified by "message" if "cond" is not nil.
	AssertIsNil(cond interface{}, message string)
	// AssertIsNotNil Reports an error identified by "message" if "cond" is nil.
	AssertIsNotNil(cond interface{}, message string)
	// TestCase
	TestCase(x interface{}, fn TestCaseFunc)
}

type goUnit struct {
	e gounitсonstraint.Evaluator
}

// New create new instance GoUnit
func New(tb testing.TB) GoUnit {
	e := gounitсonstraint.NewEvaluator(tb)
	return &goUnit{e: e}
}

func (u *goUnit) AssertTrue(cond bool, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsTrue(), cond, message)
}

func (u *goUnit) AssertFalse(cond bool, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsFalse(), cond, message)
}

func (u *goUnit) AssertError(err error, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsError(), err, message)
}

func (u *goUnit) AssertNotError(err error, message string) {
	c := gounitсonstraint.NewNot(gounitсonstraint.NewIsError())
	u.e.Evaluate(c, err, message)
}

func (u *goUnit) AssertLessThanOrEqual(expected, actual int, message string) {
	constraints := []gounitсonstraint.Constraint{
		gounitсonstraint.NewIsEquals(expected),
		gounitсonstraint.NewIsLessThan(expected),
	}
	u.e.Evaluate(gounitсonstraint.NewOr(constraints), actual, message)
}

func (u *goUnit) AssertLessThan(expected, actual int, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsLessThan(expected), actual, message)
}

func (u *goUnit) AssertGreaterThanOrEqual(expected, actual int, message string) {
	constraints := []gounitсonstraint.Constraint{
		gounitсonstraint.NewIsEquals(expected),
		gounitсonstraint.NewIsGreaterThan(expected),
	}
	u.e.Evaluate(gounitсonstraint.NewOr(constraints), actual, message)
}

func (u *goUnit) AssertGreaterThan(expected, actual int, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsGreaterThan(expected), actual, message)
}

func (u *goUnit) AssertFileEquals(expected string, actual string, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsFileEquals(expected), actual, message)
}

func (u *goUnit) AssertEquals(expected interface{}, actual interface{}, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsEquals(expected), actual, message)
}

func (u *goUnit) AssertEmpty(actual interface{}, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsEmpty(), actual, message)
}

func (u *goUnit) AssertFileExists(filename string, message string) {
	u.e.Evaluate(gounitсonstraint.NewFileExists(), filename, message)
}

func (u *goUnit) AssertDirectoryExists(dir string, message string) {
	u.e.Evaluate(gounitсonstraint.NewDirectoryExists(), dir, message)
}

func (u *goUnit) AssertCount(expectedCount int, array interface{}, message string) {
	u.e.Evaluate(gounitсonstraint.NewCount(expectedCount), array, message)
}

func (u *goUnit) AssertContains(needle interface{}, array interface{}, message string) {
	u.e.Evaluate(gounitсonstraint.NewContains(needle), array, message)
}

func (u *goUnit) AssertArrayHasKey(key interface{}, array interface{}, message string) {
	u.e.Evaluate(gounitсonstraint.NewMapHasKey(key), array, message)
}

func (u *goUnit) AssertIsNil(cond interface{}, message string) {
	u.e.Evaluate(gounitсonstraint.NewIsNil(), cond, message)
}

func (u *goUnit) AssertIsNotNil(cond interface{}, message string) {
	c := gounitсonstraint.NewNot(gounitсonstraint.NewIsNil())
	u.e.Evaluate(c, cond, message)
}

func (*goUnit) TestCase(x interface{}, fn TestCaseFunc) {
	if reflect.TypeOf(x).Kind() != reflect.Slice {
		panic("x expects is slice")
	}
	vTestCase := reflect.ValueOf(x)
	for i := 0; i < vTestCase.Len(); i++ {
		fn(vTestCase.Index(i).Interface())
	}
}
