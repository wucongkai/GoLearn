package main

import "fmt"

type User struct {
	//成员变量
	Id            int
	Score         float64
	Name, address string
}

//成员方法
func (mi User) Hello() {
	fmt.Println("My name is", mi.Name)
}

func main3() {
	var u User
	u = User{Id: 32, Score: 32.9, Name: "daqiaoqiao", address: "中国"}
	fmt.Println(u.Name)
	u.Hello()

	//匿名结构体（仅使用一次）
	var student struct {
		Name string
		Age  int
	}
	student.Name = "zcy"
	student.Age = 18

	fmt.Printf("%s, %d", student.Name, student.Age)

}
