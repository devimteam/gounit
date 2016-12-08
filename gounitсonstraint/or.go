package gounitсonstraint

type or struct {
    constraints []Constraint
}

func NewOr(constraints []Constraint) *or {
    return &or{constraints: constraints}
}

func (c *or) Matches(other interface{}) bool {
    for _, con := range c.constraints {
        if con.Matches(other) {
            return true
        }
    }
    return false
}

func (*or) Error(other interface{}) error {
    return nil
}

func (c *or) String() string {
    text := ""
    for key, con := range c.constraints {
        if key > 0 {
            text += " or "
        }
        text += con.String()
    }
    return text;
}
