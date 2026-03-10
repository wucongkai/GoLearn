package main

import "fmt"

func main5() {
	type User struct {
		Name string
		Age  int
	}

	type Vedio struct {
		Length int
		Name   string
		User   //匿名成员
	}

	u := User{Name: "zcy", Age: 18}
	v := Vedio{Length: 120, Name: "go语言教程", User: u}

	fmt.Println(v.Length)
	fmt.Println(v.Name)      //访问自己的Name
	fmt.Println(v.User.Name) //访问父类的Name
	fmt.Println(v.Age)       //vedio从User里继承了Age
}
