package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (this *Config) initForbidden()  {
	this.forbidden = make(map[string]bool)
	this.forbidden["add"] = true
	this.forbidden["add!"] = true
	this.forbidden["rm"] = true
	this.forbidden["list"] = true
	this.forbidden["help"] = true
	this.forbidden["pwd"] = true
	this.forbidden["clear!"] = true
}

func (this *Config) initDir() bool  {
	this.path = getConfigPath()
	f, err := os.Stat(this.path)
	if os.IsNotExist(err) {
		mkErr := os.Mkdir(this.path, 0755)
		check(mkErr)
		return true
	}
	if !f.IsDir() {
		fmt.Printf("已存在同名文件%s，无法创建配置文件夹\n", this.path)
		os.Exit(1)
	}
	return true
}

func (this *Config) initFile() bool  {
	this.file = getConfigFile()
	f, err := os.Stat(this.file)
	if os.IsNotExist(err) {
		_, thErr := os.Create(this.file)
		check(thErr)
		return true
	}
	if f.IsDir() {
		fmt.Printf("已存在与配置文件同名文件夹 %s，请先删除\n", this.file)
	}
	return true
}

func (this *Config) initConfig() bool {
	this.config = make(map[string]string)
	s, err := ioutil.ReadFile(this.file)
	check(err)
	settings := string(s)
	settingArr := strings.Split(settings, "\n");
	for _, line := range settingArr {
        c := strings.Split(line, "::")
		if len(c) < 2 {
			continue
		}
		this.config[c[0]] = c[1]
	}
	return true
}

func (this *Config) Init()  {
	this.initDir()
	this.initFile()
	this.initConfig()
	this.initForbidden()
	flag.Parse()
}

