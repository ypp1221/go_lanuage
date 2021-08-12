package main

import (
	"math"
	"sync"
)

type stack struct {
	array []int
	size  int
	lock  sync.Mutex
}

func (sta *stack) Push(element int) {
	sta.lock.Lock()
	defer sta.lock.Unlock()
	sta.array = append(sta.array, element)
	sta.size++
}
func (sta *stack) Empty() int {
	if sta.size == 0 {
		return 1
	}
	return 0
}
func (sta *stack) Pop() (element int) {
	if sta.Empty() == 1 {
		return math.MaxInt8
	}
	element = sta.array[sta.size-1]
	sta.size--
	return element
}
