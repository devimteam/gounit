package gounitсonstraint

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
    "time"
)

type Evaluator interface {
	Evaluate(other interface{}, description string)
}

type evaluator struct {
	c  Constraint
	tb testing.TB
}

func NewEvaluator(tb testing.TB, c Constraint) Evaluator {
	return &evaluator{c: c, tb: tb}
}

func (e *evaluator) Evaluate(other interface{}, description string) {
	_, file, line, _ := runtime.Caller(2)

    begin := time.Now()
    ok := !e.c.Matches(other)
    leadLime := time.Now().Sub(begin).Seconds()
	if ok {
		e.fail(other, file, line, leadLime, description)
	} else {
		e.success(file, line, leadLime, description)
	}
}
func (e *evaluator) getError(other interface{}) error {
	return errors.New(fmt.Sprintf("%#v", other) + " " + e.c.String())
}

func (e *evaluator) success(file string, line int, leadLime float64, description string) {
	successMsg := fmt.Sprintf("Success asserting that %s.", e.c.String())
    successMsg = e.rightPad(successMsg, " ", 80)
	successMsg = ansi.Color(successMsg, "green:")
	if description != "" {
		successMsg = "\n" + ansi.Color(description, ":blue") + "\n" + successMsg
	}
	successMsg = successMsg + ansi.Color(fmt.Sprintf("(%s:%d)", filepath.Base(file), line), "white+b:")
    successMsg = successMsg + " " + fmt.Sprintf("%fs.", leadLime)
	fmt.Println(ansi.Color("+", "green+b:") + successMsg)
}

func (e *evaluator) fail(other interface{}, file string, line int, leadLime float64, description string) {
	err := e.c.Error(other)
	if err == nil {
		err = e.getError(other)
	}

	errorMsg := fmt.Sprintf("Failed asserting that %s.", err.Error())
    errorMsg = e.rightPad(errorMsg, " ", 80)
	errorMsg = ansi.Color(errorMsg, "red:")
	if description != "" {
		errorMsg = "\n" + ansi.Color(description, ":blue") + "\n" + errorMsg
	}
	errorMsg = errorMsg + ansi.Color(fmt.Sprintf("(%s:%d)", filepath.Base(file), line), "white+b:")
    errorMsg = errorMsg + " " + fmt.Sprintf("%fs.", leadLime)
	fmt.Println(ansi.Color("-", "red+b:") + errorMsg)
	e.tb.FailNow()
}

func (e *evaluator) rightPad(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
