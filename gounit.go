package gounit

import (
	"testing"

	"github.com/l-vitaly/gounit/gounitсonstraint"
)

const (
	EmptyMessage = ""
)

type GoUnit interface {
	AssertArrayHasKey(key interface{}, array interface{}, message string)
	AssertContains(value interface{}, array interface{}, message string)
	AssertCount(count int, array interface{}, message string)
	AssertDirectoryExists(dir string, message string)
	AssertFileExists(file string, message string)
	AssertEmpty(actual interface{}, message string)
	AssertEquals(expected interface{}, actual interface{}, message string)
	AssertTrue(cond bool, message string)
	AssertFalse(cond bool, message string)
	AssertFileEquals(expected string, actual string, message string)
	AssertGreaterThan(expected, actual int, message string)
	AssertGreaterThanOrEqual(expected, actual int, message string)
	AssertLessThan(expected, actual int, message string)
	AssertLessThanOrEqual(expected, actual int, message string)
	AssertError(err error, message string)
	AssertNotError(err error, message string)
}

type goUnit struct {
	tb testing.TB
}

func New(tb testing.TB) GoUnit {
	return &goUnit{tb: tb}
}

// AssertTrue Reports an error identified by "message" if "cond" is false.
func (u *goUnit) AssertTrue(cond bool, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsTrue()).Evaluate(cond, message)
}

// AssertFalse Reports an error identified by "message" if "cond" is true.
func (u *goUnit) AssertFalse(cond bool, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsFalse()).Evaluate(cond, message)
}

//// AssertError Reports an error identified by "message" if "err" is error type
func (u *goUnit) AssertError(err error, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsError()).
		Evaluate(err, message)
}

// AssertNotError Reports an error identified by "message" if "err" is not error type
func (u *goUnit) AssertNotError(err error, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewNot(gounitсonstraint.NewIsError())).
		Evaluate(err, message)
}

// AssertLessThanOrEqual Reports an error identified by "message" if the value of
// "actual" is not less than or equal to the value of "expected".
func (u *goUnit) AssertLessThanOrEqual(expected, actual int, message string) {
	constraints := []gounitсonstraint.Constraint{
		gounitсonstraint.NewIsEquals(expected),
		gounitсonstraint.NewIsLessThan(expected),
	}
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewOr(constraints)).
		Evaluate(actual, message)
}

// AssertLessThan Reports an error identified by "message" if the value of
// "actual" is not less than the value of "expected".
func (u *goUnit) AssertLessThan(expected, actual int, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsLessThan(expected)).
		Evaluate(actual, message)
}

// AssertGreaterThanOrEqual Reports an error identified by "message" if the
// value of "actual" is not greater than or equal to the value of "expected".
func (u *goUnit) AssertGreaterThanOrEqual(expected, actual int, message string) {
	constraints := []gounitсonstraint.Constraint{
		gounitсonstraint.NewIsEquals(expected),
		gounitсonstraint.NewIsGreaterThan(expected),
	}
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewOr(constraints)).
		Evaluate(actual, message)
}

// AssertGreaterThan Reports an error identified by "message" if the value of
// "actual" is not greater than the value of "expected".
func (u *goUnit) AssertGreaterThan(expected, actual int, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsGreaterThan(expected)).
		Evaluate(actual, message)
}

// AssertFileEquals Reports an error identified by "message: if the file specified
// by "expected" does not have the same contents as the file specified by "actual".
func (u *goUnit) AssertFileEquals(expected string, actual string, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsFileEquals(expected)).
		Evaluate(actual, message)
}

// AssertEquals Reports an error identified by "message" if the two variables
// "expected" and "actual" are not equal.
func (u *goUnit) AssertEquals(expected interface{}, actual interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsEquals(expected)).
		Evaluate(actual, message)
}

// AssertEmpty Reports an error identified by "message" if "actual" is not empty.
func (u *goUnit) AssertEmpty(actual interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsEmpty()).
		Evaluate(actual, message)
}

// AssertFileExists Reports an error identified by "message" if the file specified
// by "file" does not exist.
func (u *goUnit) AssertFileExists(filename string, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewFileExists()).
		Evaluate(filename, message)
}

// AssertDirectoryExists Reports an error identified by "message" if the
// directory specified by "dir" does not exist.
func (u *goUnit) AssertDirectoryExists(dir string, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewDirectoryExists()).
		Evaluate(dir, message)
}

// Reports an error identified by "message" if the number of elements in "array"
// is not "expectedCount".
func (u *goUnit) AssertCount(expectedCount int, array interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewCount(expectedCount)).
		Evaluate(array, message)
}

// AssertContains Reports an error identified by "message" if "needle" is not
// an element of "array".
func (u *goUnit) AssertContains(needle interface{}, array interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewContains(needle)).
		Evaluate(array, message)
}

// AssertArrayHasKey Reports an error identified by "message" if "array" does
// not have the "key".
func (u *goUnit) AssertArrayHasKey(key interface{}, array interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewMapHasKey(key)).
		Evaluate(array, message)
}
