package main

import (
	"fmt"
	"github.com/KunBetter/ABTest/core"
	"io/ioutil"
)

func main() {
	ab := &core.ABTest{}
	ab.Init()

	buf, err := ioutil.ReadFile("ABExpConfig.json")
	if err != nil {
		fmt.Print(err)
		return
	}

	ab.LoadConfig([]string{string(buf)})

	reqMap := make(map[string]string)
	reqMap["type"] = "1"
	reqMap["id"] = "5"
	reqMap["userid"] = "10"
	reqMap["layId"] = "100"

	tagMap := ab.Distribute(reqMap)

	fmt.Println(tagMap)
}
