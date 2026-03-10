package main

//常量与枚举
func main2() {
	//常量
	const (
		E  = 2.17 //常量声明时必须赋值
		PI = 3.14
	)
	// 枚举
	const (
		APPLE  = 1
		HUAWEI = 2
		VIVO   = 3
	)

	//相加
	const (
		a = iota // 0
		b        // 1
		c        // 2
		d        // 3
	)

	const (
		NOT_USE = iota             //iota = 0
		KB      = 1 << (10 * iota) //iota = 1
		MB      = 1 << (10 * iota) //iota = 2
		GB      = 1 << (10 * iota) //iota = 3
	)

	const (
		NO_USE = 1 << (10 * iota)
		KKB
		MMB
		GGB
	)

}
