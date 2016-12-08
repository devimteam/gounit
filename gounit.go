package gounit

import (
	"testing"

	"github.com/l-vitaly/gounit/gounitсonstraint"
)

const (
	EmptyMessage = ""
)

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
}

type goUnit struct {
	tb testing.TB
}

func New(tb testing.TB) GoUnit {
	return &goUnit{tb: tb}
}

func (u *goUnit) AssertTrue(cond bool, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsTrue()).Evaluate(cond, message)
}

func (u *goUnit) AssertFalse(cond bool, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsFalse()).Evaluate(cond, message)
}

func (u *goUnit) AssertError(err error, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsError()).
		Evaluate(err, message)
}

func (u *goUnit) AssertNotError(err error, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewNot(gounitсonstraint.NewIsError())).
		Evaluate(err, message)
}

func (u *goUnit) AssertLessThanOrEqual(expected, actual int, message string) {
	constraints := []gounitсonstraint.Constraint{
		gounitсonstraint.NewIsEquals(expected),
		gounitсonstraint.NewIsLessThan(expected),
	}
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewOr(constraints)).
		Evaluate(actual, message)
}

func (u *goUnit) AssertLessThan(expected, actual int, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsLessThan(expected)).
		Evaluate(actual, message)
}

func (u *goUnit) AssertGreaterThanOrEqual(expected, actual int, message string) {
	constraints := []gounitсonstraint.Constraint{
		gounitсonstraint.NewIsEquals(expected),
		gounitсonstraint.NewIsGreaterThan(expected),
	}
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewOr(constraints)).
		Evaluate(actual, message)
}

func (u *goUnit) AssertGreaterThan(expected, actual int, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsGreaterThan(expected)).
		Evaluate(actual, message)
}

func (u *goUnit) AssertFileEquals(expected string, actual string, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsFileEquals(expected)).
		Evaluate(actual, message)
}

func (u *goUnit) AssertEquals(expected interface{}, actual interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsEquals(expected)).
		Evaluate(actual, message)
}

func (u *goUnit) AssertEmpty(actual interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewIsEmpty()).
		Evaluate(actual, message)
}

func (u *goUnit) AssertFileExists(filename string, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewFileExists()).
		Evaluate(filename, message)
}

func (u *goUnit) AssertDirectoryExists(dir string, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewDirectoryExists()).
		Evaluate(dir, message)
}

func (u *goUnit) AssertCount(expectedCount int, array interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewCount(expectedCount)).
		Evaluate(array, message)
}

func (u *goUnit) AssertContains(needle interface{}, array interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewContains(needle)).
		Evaluate(array, message)
}

func (u *goUnit) AssertArrayHasKey(key interface{}, array interface{}, message string) {
	gounitсonstraint.NewEvaluator(u.tb, gounitсonstraint.NewMapHasKey(key)).
		Evaluate(array, message)
}
