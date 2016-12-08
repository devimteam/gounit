package gounitсonstraint

import (
    "fmt"
    "github.com/pkg/errors"
    "os"
)

type fileExists struct {
}

func NewFileExists() *fileExists {
    return &fileExists{}
}

func (*fileExists) Matches(other interface{}) bool {
    _, err := os.Stat(other.(string))
    if !os.IsNotExist(err) {
        return true
    }
    return false
}

func (*fileExists) Error(other interface{}) error {
    return errors.New(fmt.Sprintf("file \"%s\" exists", other))
}

func (*fileExists) String() string {
    return "file exists"
}
