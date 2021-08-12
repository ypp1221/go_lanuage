package main

import (
	"fmt"
)

func main() {
	var test string = "helloworld"
	for i := 0; i < len(test); i++ {
		fmt.Printf("%v(%c)", test[i], test[i])
	}
	fmt.Println()
	for _, i := range test {
		fmt.Printf("%v(%c)", i, i)
	}

	fmt.Println()
	//字符串修改
	byte_test := []byte(test)
	byte_test[0] = 'w'
	fmt.Println(string(byte_test))

	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
	fmt.Println(c)
	d := [...]int{2, 3, 4, 5}
	fmt.Println(len(d))

	//多维数组的遍历
	var f [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	//arr := [...][3]int{{1,2,3},{4,5,6}}
	for i, j := range f {
		for k, m := range j {
			fmt.Printf("(%v,%v)= %d ", i, k, m)
		}
		fmt.Println()
	}

}
