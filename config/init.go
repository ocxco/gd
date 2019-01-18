package config

import (
	"flag"
	"fmt"
	"os"
)

func initDir() bool  {
	dir := getConfigPath()
	f, err := os.Stat(dir)
	if os.IsNotExist(err) {
		mkErr := os.Mkdir(dir, 0755)
		check(mkErr)
		return true
	}
	if !f.IsDir() {
		fmt.Printf("已存在同名文件%s，无法创建配置文件夹\n", dir)
		os.Exit(1)
	}
	return true
}

func initFile() bool  {
	file := getConfigFile()
	f, err := os.Stat(file)
	if os.IsNotExist(err) {
		_, thErr := os.Create(file)
		check(thErr)
		return true
	}
	if f.IsDir() {
		fmt.Printf("已存在与配置文件同名文件夹 %s，请先删除\n", file)
	}
	return true
}

func initFlag()  {
	flag.Bool("f", false, "-f for force add")
	flag.Parse()
}

func Init()  {
	initDir()
	initFile()
	initFlag()
}

