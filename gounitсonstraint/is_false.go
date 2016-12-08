package gounitсonstraint

type isFalse struct {
}

func NewIsFalse() *isFalse {
    return &isFalse{}
}

func (*isFalse) Matches(other interface{}) bool {
    return other == false
}

func (*isFalse) Error(other interface{}) error {
    return nil
}

func (*isFalse) String() string {
    return "is false";
}
