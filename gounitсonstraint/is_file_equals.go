package gounitсonstraint

import (
    "io/ioutil"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
)

type isFileEquals struct {
    expectedFilename string
}

func NewIsFileEquals(expectedFilename string) *isFileEquals {
    return &isFileEquals{expectedFilename: expectedFilename}
}

func (c *isFileEquals) Matches(other interface{}) bool {
    expectedByte, _ := ioutil.ReadFile(c.expectedFilename)
    actualByte, _ := ioutil.ReadFile(other.(string))

    shaExpected := sha1.New()
    shaExpected.Write(expectedByte)

    shaActual := sha1.New()
    shaActual.Write(actualByte)

    if hex.EncodeToString(shaExpected.Sum(nil)) == hex.EncodeToString(shaActual.Sum(nil)) {
        return true
    }
    return false
}

func (*isFileEquals) Error(other interface{}) error {
    return nil
}

func (c *isFileEquals) String() string {
    return fmt.Sprintf("file is equal to %#v file", c.expectedFilename);
}
