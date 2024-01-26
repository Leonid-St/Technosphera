package main

import (
	"sync"
)

// type RoundRobinBalancer interface {
// 	GiveStat()
// 	Init()
// 	GiveNode()
// }

func (b *RoundRobinBalancer) GiveStat() []int {
	return b.stat
}

func (b *RoundRobinBalancer) Init(n int) {
	b.stat = make([]int, n)
}

func (b *RoundRobinBalancer) GiveNode() int {
	b.RWMutex.Lock()
	inc := (func() { b.i++; b.RWMutex.Unlock() })
	if b.i == len(b.stat) {
		b.i = 0
		b.stat[b.i]++
		defer inc()
		return b.i
	} else {
		b.stat[b.i]++
		defer inc()
		return b.i
	}

}

type RoundRobinBalancer struct {
	stat []int
	i    int
	sync.RWMutex
}

func main() {
	// balancer := new(RoundRobinBalancer)
	// balancer.Init(3)
	// balancer.GiveStat()

	// expected := []int{0, 0, 0}
	// result := balancer.GiveStat()
	// fmt.Println("expected", expected, "have", result)

}
