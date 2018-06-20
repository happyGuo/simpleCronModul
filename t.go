package main

import (
	"github.com/pelletier/go-toml"

	//"path/filepath"
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	workPath, _ := os.Getwd()
	filepath := filepath.Join(workPath, "zycron","app","conf")+"\\mysql.toml"
	if _,err := os.Stat(filepath);err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}

	conf, _ := toml.LoadFile(filepath)
	envirmentConf :=conf.Get("dev").(*toml.Tree)
	fmt.Println(envirmentConf.Get("host"))
}