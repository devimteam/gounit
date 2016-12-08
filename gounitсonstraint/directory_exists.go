package gounitсonstraint

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type directoryExists struct {
}

func NewDirectoryExists() *directoryExists {
	return &directoryExists{}
}

func (*directoryExists) Matches(other interface{}) bool {
	if stat, err := os.Stat(other.(string)); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func (*directoryExists) Error(other interface{}) error {
	return errors.New(fmt.Sprintf("directory \"%s\" exists", other))
}

func (*directoryExists) String() string {
	return "directory exists"
}
