package config

import (
	"fmt"
	"os"
)

func check(err error)  {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getConfigPath() string  {
	configPath := os.Getenv("HOME") + "/.config/gd/"
	return configPath
}

func getConfigFile() string  {
	home := os.Getenv("HOME")
	return home + "/.config/gd/setting"
}

type Config struct {
	file string // config file
	path string // config path
	changed bool // mark if the config had changed
	config map[string]string // parsed configs
	forbidden map[string]bool // keywords for forbidden
	params params // command line params
}

type params struct {
	add *string
	rm *string
	list *bool
	f *bool
}