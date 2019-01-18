package main

import (
	"fmt"
	"gd/config"
)

func main() {
	config.Init()
	path := config.Get("go2")
	fmt.Println(path)
}
