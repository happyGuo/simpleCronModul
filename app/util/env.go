package util

import (
	"os"
)
//GetEnv 获取环境
func GetEnv() string {

	if model := os.Getenv("RUNMODE");model!=""{
		return model
	}
	return "dev"
}

