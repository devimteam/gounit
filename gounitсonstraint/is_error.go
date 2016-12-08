package gounitсonstraint

type isError struct {
}

func NewIsError() *isError {
	return &isError{}
}

func (*isError) Matches(other interface{}) bool {
	if other != nil {
		_, ok := other.(error)
		return ok
	}
	return false
}

func (*isError) Error(other interface{}) error {
	return nil
}

func (*isError) String() string {
	return "is error"
}
