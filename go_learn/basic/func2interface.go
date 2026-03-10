package main

type Selector interface {
	Select([]string) int
}

type ConnectionPool struct {
	Servers      []string
	LoadBalancer Selector //成员变量是接口类型
}

func f1([]string) int {
	return 0
}

type RoundRobin struct{}

func (RoundRobin) Select(s []string) int { return f1(s) }

func mainAA() {
	cp := ConnectionPool{
		Servers:      []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		LoadBalancer: RoundRobin{},
	}

	cp.LoadBalancer.Select(cp.Servers)
}
