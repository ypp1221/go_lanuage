package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Print("say:", i)
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
	fmt.Println()

	//gosched使用
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	//主线程
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Print("hello")
	}

	time.Sleep(2000)

}
