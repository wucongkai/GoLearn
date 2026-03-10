package main

import (
	"fmt"
	"strconv"
)

func asign_string() {
	//字符串里可以包含任意Unicode字符
	s1 := "My name is kai"

	//包含转义字符
	s2 := "He say:\"I'm fine.\"\n\\Thank\tyou.\\"
	//反引号里的转义字符无效，

	s3 := `here is first line.
	
	there is third line.
	`

	fmt.Println("s1")
	fmt.Println(s1)
	fmt.Println("s2")
	fmt.Println(s2)
	fmt.Println("s3")
	fmt.Println(s3)
}

func string_other_convert() {
	var err error
	var i int = 8
	var i64 int64 = int64(i)

	//int转string
	var s string = strconv.Itoa(i) //内部调用FormatInt
	s = strconv.FormatInt(i64, 10)

	//string 转int
	i, err = strconv.Atoi(s)

	//string转int64
	i64, err = strconv.ParseInt(s, 10, 64)

	//float转string
	var f float64 = 8.123456789
	s = strconv.FormatFloat(f, 'f', 2, 64) //	保留2位小数 %2f
	fmt.Println(s)

	//string转float
	_ = err
}
