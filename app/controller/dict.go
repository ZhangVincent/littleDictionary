package controller

// @description 字典方法
// @author zkp15
// @date 2023/7/25 20:55
// @version 1.0

import (
	"fmt"
	"littleDictionary/app/model"
	"littleDictionary/app/service"
	"log"
	"strings"
)

func Dict() {
	fmt.Println("=========================")
	fmt.Println("请输入待翻译的值，例如：\n0 你好\n1 hello")
	var (
		types  int    = 0
		source string = "你好"
	)
	fmt.Scanf("%d%s", &types, &source)
	if types != 0 && types != 1 {
		log.Fatal("只支持两种翻译格式！")
	}

	//if len(os.Args) !=3{
	//	log.Fatal("输入格式错误")
	//}
	//types := os.Args[1]
	//source := os.Args[2]

	transtype := []string{"zh2en", "en2zh"}
	dictRequest := model.DictRequest{
		TransType: transtype[types],
		Source:    source,
	}

	dictResponse := service.Query(dictRequest)

	if types == 0 {
		dictionary := dictResponse.Dictionary.Explanations[0]
		fmt.Println(dictionary[strings.LastIndex(dictionary, ";")+2:])
		fmt.Println(dictResponse.Dictionary.Explanations[0])
	} else {
		prons := dictResponse.Dictionary.Prons
		fmt.Printf("%v  %v  %v\n", dictResponse.Dictionary.Entry, prons.En, prons.EnUs)
		explanations := dictResponse.Dictionary.Explanations
		for _, s := range explanations {
			fmt.Println(s)
		}
		example := dictResponse.Dictionary.WqxExample
		for i := 0; i < len(example); i++ {
			fmt.Printf("%v  %v\n", example[i][0], example[i][1])
		}
	}
}
