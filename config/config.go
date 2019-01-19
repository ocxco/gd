package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (this *Config) checkCanAdd(name string, path string)  {
	if name == "" {
		fmt.Println("Must assign a name")
		os.Exit(1)
	}
	_, ok := this.forbidden[name]
	if ok {
		fmt.Println("Can not use add、rm、list as short keyword")
		os.Exit(1)
	}
	for key, value := range this.config {
		if path == value {
			fmt.Printf("Current directory %s already added as %s\n", value, key)
			os.Exit(1)
		}
	}
}

func (this *Config) remove(name string)  {
	if name == "" {
		fmt.Println("Must assign a name")
		os.Exit(1)
	}
	delete(this.config, name);
	fmt.Printf("Success, %s is removed from config\n", name)
	this.changed = true
}

func (this *Config) add(name string, path string)  {
	this.checkCanAdd(name, path)
	nowPath := this.get(name)
	if "" != nowPath {
		fmt.Printf("%s already added, now path is %s\n", name, nowPath)
		os.Exit(1)
	} else {
		this.config[name] = path
		fmt.Printf("Success, %s is added to config as '%s' alias\n", name, path)
	}
	this.changed = true
}

func (this *Config) forceAdd(name string, path string)() {
	this.checkCanAdd(name, path)
	this.config[name] = path
	fmt.Printf("Success, %s is added to config as '%s' alias\n", name, path)
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
	fmt.Println("Usage: gd [command] [name]")
	fmt.Printf("  %-12s %s\n", "add  <name>", "add current directory as name")
	fmt.Printf("  %-12s %s\n", "add! <name>", "force add current directory as name")
	fmt.Printf("  %-12s %s\n", "ls | list", "list all of configs")
	fmt.Printf("  %-12s %s\n", "pwd", "show current directory")
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
			fmt.Printf("%s is not set\n", flag.Arg(0))
			os.Exit(1)
		}
		fmt.Println(path)
		os.Exit(0)
	}
	this.save()
	os.Exit(1)
}