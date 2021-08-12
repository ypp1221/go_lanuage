package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(arr [10]int) (sum int) {
	for _, i := range arr {
		sum += i
	}
	return sum
}

func find_index(arr [10]int, target int) {
	for i := 0; i < len(arr); i++ {
		value := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if value == arr[j] {
				fmt.Printf("index:(%v, %v)", i, j)
			}
		}
	}
}
func main() {
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	sum_value := sum(b)
	fmt.Println(sum_value)

	var two_sum [10]int = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	find_index(two_sum, 10)

}
