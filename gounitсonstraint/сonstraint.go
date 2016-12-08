package gounitсonstraint

type Constraint interface {
	Matches(other interface{}) bool
    Error(other interface{}) error
    String() string
}
