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
