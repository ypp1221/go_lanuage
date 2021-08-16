package main

import (
	"fmt"
	"sync"
)

type singleinstance struct {
	Name string
}

var single *singleinstance
var lock sync.Mutex

func GetSingleObject() *singleinstance {
	lock.Lock()
	defer lock.Unlock()
	if single == nil {
		fmt.Println("只能执行一次")
		single = &singleinstance{"yang"}
	}
	return single
}
func (single *singleinstance) PringName() {
	fmt.Println("唯一的名字是:", single.Name)
}
func main() {
	a := GetSingleObject()
	a.PringName()
	b := GetSingleObject()
	b.PringName()
	if a == b {
		fmt.Println("a same as b")
	}
	fmt.Println(&a.Name, &b.Name, &single.Name)
}
