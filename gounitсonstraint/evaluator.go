package gounitсonstraint

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"time"

	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
)

type Evaluator interface {
	Evaluate(c Constraint, other interface{}, description string)
}

type evaluator struct {
	tb testing.TB
}

func NewEvaluator(tb testing.TB) Evaluator {
	return &evaluator{tb: tb}
}

func (e *evaluator) Evaluate(c Constraint, other interface{}, description string) {
	_, file, line, _ := runtime.Caller(2)

	begin := time.Now()
	result := c.Matches(other)
	leadLime := int(time.Now().Sub(begin).Nanoseconds() / 1000)

	if !result {
		err := c.Error(other)
		if err == nil {
			err = e.getError(c, other)
		}

		msg := e.makeMessage(true, file, line, leadLime, err.Error(), description)
		fmt.Println(msg)
		e.tb.FailNow()
	} else {
		msg := e.makeMessage(false, file, line, leadLime, c.String(), description)
		fmt.Println(msg)
	}
}
func (e *evaluator) getError(c Constraint, other interface{}) error {
	return errors.New(fmt.Sprintf("%#v", other) + " " + c.String())
}

func (e *evaluator) makeMessage(isError bool, file string, line int, leadLime int, message string, description string) string {
	title := "Success"
	color := "green"
	if isError {
		title = "Failed"
		color = "red"
	}
	msg := ansi.Color(fmt.Sprintf("%s asserting that %s", title, message), color+":")

	if description != "" {
		msg = msg + fmt.Sprintf(" (%s)", description)
	}

	if isError {
		msg = msg + fmt.Sprintf(" [%s:%d]", filepath.Base(file), line)
	}

	msg = msg + fmt.Sprintf(" %d ms", leadLime)

	return msg
}

func (e *evaluator) rightPad(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
