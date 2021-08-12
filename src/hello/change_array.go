package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type change_array struct {
	array []int
	len   int
	cap   int
	lock  sync.Mutex
}

func Make_array(len int, cap int) *change_array {
	s := new(change_array)
	if len > cap {
		panic("len > cap, this is error")
	}
	slice_array := make([]int, cap, cap)
	s.array = slice_array
	s.len = 0
	s.cap = cap
	return s

}

func (s *change_array) Add_element(element int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == s.cap {
		//没有容量了，要进行扩容
		new_cap := 2 * s.cap
		if s.cap == 0 {
			new_cap = 1
		}
		new_array := make([]int, new_cap, new_cap)
		//复制s.array到new_array
		for i, v := range s.array {
			new_array[i] = v
		}
		s.array = new_array
		s.cap = new_cap
	}
	//将加入的元素放入新数组的最后一个位置
	s.array[s.len] = element
	s.len++
}
func (s *change_array) Print_array() {
	for i := 0; i < s.len; i++ {
		fmt.Printf("%d ", s.array[i])
	}
	fmt.Println()
}

func (s *change_array) Add_most_element() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		s.Add_element(rand.Intn(100))
	}
	s.Print_array()
}
func main() {
	slice_array := Make_array(0, 0)
	slice_array.Print_array()
	slice_array.Add_element(4)
	slice_array.Print_array()
	slice_array.Add_most_element()
}
