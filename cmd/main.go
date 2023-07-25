package main

import (
	"fmt"
	"littleDictionary/app/controller"
	"log"
)

// @description 唯一接口
// @author zkp15
// @date 2023/7/25 19:23
// @version 1.0
func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			log.Println(err)
		}
	}()
	for {
		controller.Dict()
	}
}
