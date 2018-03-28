package main

import (
	"fmt"

	"github.com/phongntt/go-spider-monitor/config"
)

func main() {
	confFile := "./conf/config.json"

	config, err := config.ReadFromFile(confFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
}
