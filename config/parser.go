package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parse() []string {
	file := getConfigFile()
	s, err := ioutil.ReadFile(file)
	check(err)
	settings := string(s)
	settingArr := strings.Split(settings, "\n");
	return settingArr
}

func Get(name string) string {
	settings := parse()
	for i := 0; i < len(settings); i++  {
		setting := strings.Split(settings[i], ":")
		if setting[0] == name {
			return setting[1]
		}
	}
	return ""
}

func add(name string, path string)  {
	setting := name + ":" + path + "\n"
	b := []byte(setting)
	err := ioutil.WriteFile(getConfigFile(), b, 0755)
	check(err)
}

func forceAdd()  {
	settings := parse()
}

func Add(name string, path string)  {
	nowPath := Get(name)
	if "" != nowPath {
		fmt.Printf("%s already added, now path is %s\n", name, nowPath)
		os.Exit(1)
	}
}