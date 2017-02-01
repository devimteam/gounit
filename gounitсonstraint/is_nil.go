package gounitсonstraint

type isNil struct {
}

func NewIsNil() *isNil {
	return &isNil{}
}

func (*isNil) Matches(other interface{}) bool {
	return other == nil
}

func (*isNil) Error(other interface{}) error {
	return nil
}

func (*isNil) String() string {
	return "is nil"
}
