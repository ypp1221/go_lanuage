package main

import (
	"container/list"
	"fmt"
)

func main() {
	s := list.New()
	for i := 1; i < 10; i++ {
		s.PushFront(i)
	}
	fmt.Println(s.Len())
	fmt.Println(s.Front())
	for e := s.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
