package main

import "fmt"

func array1d() {
	//数组必须指定长度和类型，且长度和类型指定后不可改变
	var arr1 [5]int = [5]int{}
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}            //给前两个元素赋值
	var arr4 = [5]int{2: 15, 4: 30}    //指定index赋值
	var arr5 = [...]int{3, 2, 6, 5, 4} //根据{}里元素的个数，推断出数组的长度
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}} //数组的元素类型有由匿名结构体给定

	fmt.Printf("arr1 = %#v\n", arr1)
	fmt.Printf("arr2 = %#v\n", arr2)
	fmt.Printf("arr3 = %#v\n", arr3)
	fmt.Printf("arr4 = %#v\n", arr4)
	fmt.Printf("arr5 = %#v\n", arr5)
	fmt.Printf("arr6 = %#v\n", arr6)

	fmt.Println("*******&***")

	//通过index访问数组里的元素
	fmt.Printf("arr5[0] = %d \n", arr5[0])
	fmt.Printf("arr5[len(arr5) - 1] = %d\n", arr5[len(arr5)-1])
	fmt.Printf("数组的地址： %p\n", &arr5)
	fmt.Printf("第一个元素的地址：%p\n", &arr5[0])
	fmt.Printf("第二个元素的地址：%p\n", &arr5[1])

	fmt.Println("*******&***")

	//遍历数组里的元素
	for i, ele := range arr5 {
		fmt.Printf("index = %d, element = %d \n", i, ele)
	}

	//另外一种遍历方式
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("index = %d, element = %d \n", i, arr5[i])
	}

	//数组的长度是不变的，capacity和length相等
	fmt.Printf("len(arr1) = %d\n", len(arr1))
	fmt.Printf("cap(arr1) = %d\n", cap(arr1))
}

func array2d() {
	//五行三列，只给前两行赋值，且前两行的所有列还没有赋满
	var arr1 = [5][3]int{{1}, {2, 3}}
	//第一维 可以用...推测，第二维不能用...
	var arr2 = [...][3]int{{1}, {1, 2, 3}}
	//根据index访问数组里的元素
	fmt.Printf("arr[1][1] = %d\n", arr1[1][1])
	fmt.Printf("arr[4][2] = %d\n", arr1[4][2]) //最后一个元素

	//遍历二维数组
	for row, array := range arr2 { //先取出某一行
		for col, ele := range array { //再遍历这一行
			fmt.Printf("arr2[%d][%d] = %d\n", row, col, ele)
		}
	}

	//对于多维数组，其cap和len指第一维的长度
	fmt.Printf("len(arr1) = %d\n", len(arr1))
	fmt.Printf("cap(arr1) = %d\n", cap(arr1))
}

//参数必须是长度为5的int型数组（注意长度必须是5）
//
//go语言中函数传参都是按值传递，即传递数组实际上传的是数组的拷贝
func update_array1(arr [5]int) {
	fmt.Printf("array in function, address is %p\n", &arr[0])
	arr[0] = 888
}

//传数组的指针
func update_array2(arr *[5]int) {
	fmt.Printf("array in function, asddress is %p\n", &((*arr)[0]))
	arr[0] = 888 //因为传的是数组指针，所以直接在原来的内存空间上进行修改
}

func update_array3(arr [5]*int) {
	*arr[0] = 888 //因为传的是int指针，所以直接在原来的内存空间上进行修改
}

func for_range_array() {
	arr := [...]int{1, 2, 3}

	for i, ele := range arr {
		arr[i] += 8 //修改arr里的元素，不影响ele
		fmt.Printf("%d %d %d \n", i, arr[i], ele)
		ele += 1 //修改ele不影响arr
		fmt.Printf("%d %d %d \n", i, arr[i], ele)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
}

func mail() {
	array1d()
	array2d()
}
