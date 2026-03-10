package main

func main22() {
	var f1, f2 func(a, b int, c string, d bool) (int, bool)

	f1 = func(a, b int, c string, d bool) (int, bool) {
		return a + b, d && c == "abc"
	}
	i, p := f1(1, 2, "3", false)
	_, _ = i, p

	type ConnectionPool struct {
		Servers      []string
		LoadBalancer func(a, b int, c string, d bool) (int, bool) //成员变量是接口类型
	}

	cp := ConnectionPool{
		Servers:      []string{"127.0.0.1:1234", "127.0.0.1:5768"},
		LoadBalancer: f1,
	}

	_ = cp
	_ = f2

	cp.LoadBalancer(1, 2, "3", false)
}
