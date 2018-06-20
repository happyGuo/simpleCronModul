package util

import (
	"testing"
	"fmt"
)

func TestSetConfig(t *testing.T) {
	fmt.Println(GetConfig("mysql","host"))
	fmt.Println(GetEnv())

}
