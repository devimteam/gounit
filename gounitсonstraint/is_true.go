package gounitсonstraint

type isTrue struct {
}

func NewIsTrue() *isTrue {
    return &isTrue{}
}

func (*isTrue) Matches(other interface{}) bool {
    return other == true
}

func (*isTrue) Error(other interface{}) error {
    return nil
}

func (*isTrue) String() string {
    return "is true";
}
