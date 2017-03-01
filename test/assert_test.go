package test

import (
	"errors"
	"testing"

	"github.com/l-vitaly/gounit"
)

type testStruct struct {
	Field1 string
	Field2 int
}

func TestAssertArrayHasKey(t *testing.T) {
	u := gounit.New(t)
	testArray := map[string]int{"foo": 0}
	u.AssertArrayHasKey("foo", testArray, gounit.EmptyMessage)
}

func TestAssertContains(t *testing.T) {
	u := gounit.New(t)
	testArray := []int{34, 45, 5, 22}
	u.AssertContains(45, testArray, gounit.EmptyMessage)
	u.AssertContains(5, testArray, gounit.EmptyMessage)
}

func TestAssertCount(t *testing.T) {
	u := gounit.New(t)
	testArray := []int{34, 45, 5, 22}
	u.AssertCount(4, testArray, gounit.EmptyMessage)
}

func TestAssertDirectoryExists(t *testing.T) {
	u := gounit.New(t)
	u.AssertDirectoryExists("/etc", gounit.EmptyMessage)
}

func TestAssertFileExists(t *testing.T) {
	u := gounit.New(t)
	u.AssertFileExists("./expected.txt", gounit.EmptyMessage)
}

func TestAssertEmpty(t *testing.T) {
	var mapTest map[string]int
	u := gounit.New(t)

	u.AssertEmpty(0, gounit.EmptyMessage)
	u.AssertEmpty(0.0, gounit.EmptyMessage)
	u.AssertEmpty(false, gounit.EmptyMessage)
	u.AssertEmpty("", gounit.EmptyMessage)
	u.AssertEmpty(testStruct{}, gounit.EmptyMessage)
	u.AssertEmpty(mapTest, gounit.EmptyMessage)
}

func TestAssertEquals(t *testing.T) {
	u := gounit.New(t)
	u.AssertEquals("", "", gounit.EmptyMessage)
	u.AssertEquals(0, 0, gounit.EmptyMessage)
	u.AssertEquals(true, true, gounit.EmptyMessage)
}

func TestAssertFileEquals(t *testing.T) {
	u := gounit.New(t)
	u.AssertFileEquals("./expected.txt", "./actual.txt", gounit.EmptyMessage)
}

func TestAssertGreaterThan(t *testing.T) {
	u := gounit.New(t)
	u.AssertGreaterThan(2, 4, gounit.EmptyMessage)
}

func TestAssertGreaterThanOrEqual(t *testing.T) {
	u := gounit.New(t)
	u.AssertGreaterThanOrEqual(5, 8, gounit.EmptyMessage)
	u.AssertGreaterThanOrEqual(3, 3, gounit.EmptyMessage)
}

func TestAssertLessThan(t *testing.T) {
	u := gounit.New(t)
	u.AssertLessThan(2, 1, gounit.EmptyMessage)
}

func TestAssertLessThanOrEqual(t *testing.T) {
	u := gounit.New(t)
	u.AssertLessThanOrEqual(8, 5, gounit.EmptyMessage)
	u.AssertLessThanOrEqual(3, 2, gounit.EmptyMessage)
}

func TestAssertError(t *testing.T) {
	u := gounit.New(t)
	u.AssertError(errors.New("test error"), gounit.EmptyMessage)
}

func TestAssertNotError(t *testing.T) {
	u := gounit.New(t)
	u.AssertNotError(nil, gounit.EmptyMessage)
}

func TestAssertNot(t *testing.T) {
	u := gounit.New(t)
	u.AssertNotNil("", gounit.EmptyMessage)
	u.AssertNotNil(0, gounit.EmptyMessage)
	u.AssertNotNil(false, gounit.EmptyMessage)
}
