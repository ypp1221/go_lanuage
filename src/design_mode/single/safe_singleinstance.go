package main

import (
	"fmt"
	"sync"
)

type safesingle struct {
}

var safeobject *safesingle
var once sync.Once
var wg sync.WaitGroup

func GetSafeObject() *safesingle {
	once.Do(func() {
		safeobject = &safesingle{}
		fmt.Println("第一次创建单例")
	})
	wg.Done()
	return safeobject
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go GetSafeObject()
	}
	wg.Wait()
}
