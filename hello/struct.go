package main

import (
	"fmt"
)

type person struct {
	name string
	city string
	age  int
}

//构造函数
func build_person(name string, city string, age int) person {
	return person{
		name: name,
		city: city,
		age:  age,
	}
}

func demo(ce []person) {
	ce[0].name = "modify"
}

//方法属于特定类型，函数不属于
func (p person) Dream() {
	fmt.Println("value:you are dreaming")
}
func (p *person) Swimming() {
	fmt.Println("point: you are swimming")
}
func main() {
	p := build_person("zhang", "zhengzhou", 22)
	p.Dream()
	p.Swimming()

	p1 := build_person("zhangs", "zhengzhous", 22)
	p2 := &p1
	p2.Dream()
	p2.Swimming()
	ce := []person{{"zhang", "beijing", 22}, {"yang", "shanghai", 44}}
	demo(ce)
	fmt.Println(ce)
}
