package gounitсonstraint

type isError struct {
}

func NewIsError() *isError {
	return &isError{}
}

func (*isError) Matches(other interface{}) bool {
    _, ok := other.(error)
    return ok
}

func (*isError) Error(other interface{}) error {
	return nil
}

func (*isError) String() string {
	return "is error"
}
