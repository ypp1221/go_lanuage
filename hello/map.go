package main

import (
	"fmt"
	"sort"
)

func main() {
	var person map[string]int = make(map[string]int, 10)
	person["yang"] = 10
	person["zhang"] = 12
	fmt.Println(person)
	fmt.Printf("%T\n", person)

	//判断某个键值是否存在
	value, ok := person["zhang"]
	if ok {
		fmt.Printf("value is %v\n", value)
	} else {
		fmt.Println("value is empty")
	}
	//map的遍历
	for k, value := range person {
		fmt.Printf("%s,%v:", k, value)
	}

	//按照指定顺序遍历map
	MAX := 100
	fmt.Println()
	sorce := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		sorce[key] = i + MAX
	}
	for ke, va := range sorce {
		fmt.Printf("(%s, %d)", ke, va)
	}
	fmt.Println()
	keys := make([]string, 0, 100)
	for key, _ := range sorce {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println()
	for _, k := range keys {
		fmt.Printf("%d ", sorce[k])
	}

}
