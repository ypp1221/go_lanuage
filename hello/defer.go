package main

import "fmt"

func main() {
	var test_defer [5]struct{}
	for i := 0; i < len(test_defer); i++ {
		defer fmt.Println(i)
	}
}
