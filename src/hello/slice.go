package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	p := &s[2]
	*p += 100
	fmt.Println(s)

	//使用make进行slice的创建, len(sl) = 6, cap(sl) = 10
	sl := make([]int, 6, 10)
	fmt.Println(sl)

	//slice中的append方法的使用

	test_slice_a := make([]int, 2, 4)
	test_slice_b := []int{3, 4, 5}

	//append返回的是新的slice对象，因此在内存上存在差异
	test_slice_c := append(test_slice_a, 4, 5, 6, 6)
	fmt.Printf("test_slice_a: %p, test_slice_c:%p\n", test_slice_a, test_slice_c)
	test_slice_aAndb := append(test_slice_a, test_slice_b...)
	fmt.Println("test_slice_aAndb:", test_slice_aAndb)
	fmt.Println("test_slice_c:", test_slice_c)

	//指针的测试
	var a int
	fmt.Printf("%p", &a)
	fmt.Println()
	var ptr *int = &a
	*ptr = 100
	fmt.Println(a)
}
