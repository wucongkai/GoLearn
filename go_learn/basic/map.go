package main

import "fmt"

func updata_map() {
	var m map[string]int        //声明
	m = make(map[string]int)    //初始化，容量为0
	m = make(map[string]int, 5) //初始化，容量为5.强烈建议初始化时给一个合适的容量，减少扩容的概率

	//初始化时直接赋值
	m = map[string]int{
		"语文": 0,
		"数学": 39,
		"物理": 57,
		"历史": 49,
		"科学": 110,
	}

	//往map里添加key-value对
	m["英语"] = 58
	fmt.Println(m["数学"])
	delete(m, "数学")

	if value, exist := m["语文"]; exist {
		fmt.Println(value)
	} else {
		fmt.Println("map里不存在语文这个key")
	}

	//获取map的长度，无法获取map的cap
	fmt.Printf("map里有%d对KV\n", len(m))

	//遍历map
	for key, value := range m {
		fmt.Printf("%s= %d\n", key, value)
	}

	fmt.Println("---------")
}
