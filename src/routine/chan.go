package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("recv success:", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 2
	fmt.Println("send success")
	close(ch)
}
