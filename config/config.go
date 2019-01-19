package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (this *Config) Get(name string) string {
	cfg, ok := this.config[name]
	if ok {
		return cfg
	}
	return ""
}

func (this *Config) isForbidden(name string)  {
	_, ok := this.forbidden[name]
	if ok {
		fmt.Println("Can not use add、rm、list as short keyword\n")
		os.Exit(2)
	}
}

func (this *Config) remove(name string)  {
	delete(this.config, name);
	this.changed = true
}

func (this *Config) add(name string)  {
	this.isForbidden(name)
	nowPath := this.Get(name)
	dir, err := os.Getwd()
	check(err)
	if "" != nowPath {
		if !*this.params.f {
			fmt.Printf("%s already added, now path is %s\n", name, nowPath)
			os.Exit(1)
		}
		this.config[name] = dir
	} else {
		this.config[name] = dir
	}
	this.changed = true
}

func (this *Config) list()  {
	for key, value := range this.config {
		fmt.Printf("%s => %s\n", key, value)
	}
}

func (this *Config) save()  {
	if !this.changed {
		return
	}
	err := os.Remove(this.file)
	check(err)
	settingArr := []string{}
	for key, value := range this.config {
		settingArr = append(settingArr, key + ":" + value)
	}
	settings := strings.Join(settingArr, "\n") + "\n"
	b := []byte(settings)
	err = ioutil.WriteFile(this.file, b, 0666)
	check(err)
}

func (this *Config) Startup()  {
	this.Init()
	if *this.params.add != "" {
		this.add(*this.params.add)
	} else if *this.params.rm != "" {
		this.remove(*this.params.rm)
	} else if *this.params.list {
		this.list()
	} else {
		path := this.Get(flag.Arg(0))
		if path != "" {
			err := os.Chdir(path)
			check(err)
		}
	}
	this.save()
}