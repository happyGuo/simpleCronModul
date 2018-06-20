package util

import (
	"path/filepath"
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)


//setConfig  加载配置文件
func GetConfig(confFileName,key string) string{

	filepath := filepath.Join(getConfFilePath(),confFileName+".toml")
	if _,err := os.Stat(filepath);err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}

	conf, _ := toml.LoadFile(filepath)

	envirmentConf :=conf.Get(GetEnv()).(*toml.Tree)

	return envirmentConf.Get(key).(string)

}
//getConfFilePath  获取配置文件目录
func getConfFilePath() string{

	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	appConfigPath := filepath.Join(workPath,"app","conf")

	if !FileExists(appConfigPath) {
		appConfigPath = filepath.Join(appPath, "app","conf")
	}

	return appConfigPath

}