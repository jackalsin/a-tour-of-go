package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCount struct {
	c sync.Mutex
	m map[string]int
}

func (safeCount *SafeCount) Inc(key string) {
	safeCount.c.Lock()
	safeCount.m[key]++
	safeCount.c.Unlock()
}

func (safeCount *SafeCount) Get(key string) int {
	safeCount.c.Lock()
	defer safeCount.c.Unlock()
	return safeCount.m[key]
}

func main() {
	const key = "key1"
	var safeCount = SafeCount{m: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go safeCount.Inc(key)
	}
	time.Sleep(time.Second)
	fmt.Println(safeCount.Get(key))
}
