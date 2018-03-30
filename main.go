package main

import (
	"fmt"

	"github.com/phongntt/go-spider-monitor/config"
	"github.com/phongntt/go-spider-monitor/spiderutils"
)

func main() {
	confFile := "./conf/config.json"

	config, err := config.ReadFromFile(confFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)

	text := "My name is Nguyen Tran Tuan Phong"
	enText, errEn := spiderutils.SpiderEncrypt(text)
	if errEn != nil {
		fmt.Println(errEn)
	}
	fmt.Println("Encrypted --> ", enText)
	deText, errDe := spiderutils.SpiderDecrypt(enText)
	if errDe != nil {
		fmt.Println(errDe)
	}
	fmt.Println("Decrypted --> ", deText)
}
