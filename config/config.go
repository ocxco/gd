package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func (this *Config) checkCanAdd(name string, path string)  {
	if name == "" {
		fmt.Println("Must assign an alias to add")
		os.Exit(1)
	}
	_, ok := this.forbidden[name]
	if ok {
		fmt.Println("Can not use add、add!、rm、list、pwd、clear! as alias")
		os.Exit(1)
	}
	for key, value := range this.config {
		if path == value {
			if flag.Arg(0) == "add!" {
				delete(this.config, key)
				return
			}
			fmt.Printf("Current directory %s already added as alias %s, use add! to realias it\n", value, key)
			os.Exit(1)
		}
	}
}

func (this *Config) remove(name string)  {
	if name == "" {
		fmt.Println("Must assign an alias to remove")
		os.Exit(1)
	}
	delete(this.config, name);
	fmt.Printf("Success, alias %s is removed from config\n", name)
	this.changed = true
}

func (this *Config) add(name string, path string)  {
	this.checkCanAdd(name, path)
	nowPath := this.get(name)
	if "" != nowPath {
		fmt.Printf("alias %s already used, now path is %s\n", name, nowPath)
		os.Exit(1)
	} else {
		this.config[name] = path
		fmt.Printf("Success,alias %s for %s is added to config \n", name, path)
	}
	this.changed = true
}

func (this *Config) forceAdd(name string, path string)() {
	this.checkCanAdd(name, path)
	this.config[name] = path
	fmt.Printf("Success,alias %s for %s is added to config \n", name, path)
	this.changed = true
}

func (this *Config) list()  {
	maxLen := 0.0
	for key, _ := range this.config {
		maxLen = math.Max(maxLen, float64(len(key)))
	}
	for key, value := range this.config {
		fmt.Printf("%[1]*s => %s\n", int(maxLen), key, value)
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
        settingArr = append(settingArr, key + "::" + value)
	}
	settings := strings.Join(settingArr, "\n") + "\n"
	b := []byte(settings)
	err = ioutil.WriteFile(this.file, b, 0666)
	check(err)
}

func (this *Config) get(name string) string {
	cfg, ok := this.config[name]
	if ok {
		return cfg
	}
	return ""
}

func (this *Config) help() {
	fmt.Println("Usage: gd [command] [alias]")
	fmt.Printf("  %-12s %s\n", "add  <alias>", "add current directory as alias")
	fmt.Printf("  %-12s %s\n", "add! <alias>", "force add current directory as alias")
	fmt.Printf("  %-12s %s\n", "rm   <alias>", "remove an exist alias from list")
	fmt.Printf("  %-12s %s\n", "ls | list", "list all of configs")
	fmt.Printf("  %-12s %s\n", "pwd", "show current directory, it's useful on git bash on windows")
	fmt.Printf("  %-12s %s\n", "help", "show this help content")
	fmt.Printf("  %-12s %s\n", "clear!", "clear all config")
}

func (this *Config) Startup() {
	this.Init()
	switch flag.Arg(0) {
	case "add":
		dir, err := os.Getwd()
		check(err)
		this.add(flag.Arg(1), dir)
	case "add!":
		dir, err := os.Getwd()
		check(err)
		this.forceAdd(flag.Arg(1), dir)
	case "rm":
		this.remove(flag.Arg(1))
	case "ls", "list":
		this.list()
	case "", "help":
		this.help()
	case "pwd":
		dir, err := os.Getwd()
		check(err)
		fmt.Println(dir)
	case "clear!":
		this.config = map[string]string{}
		this.changed = true
		fmt.Println("all config cleared")
	default:
		path := this.get(flag.Arg(0))
		if path == "" {
			fmt.Printf("alias %s haven't set\n", flag.Arg(0))
			os.Exit(1)
		}
		fmt.Println(path)
		os.Exit(0)
	}
	this.save()
	os.Exit(1)
}