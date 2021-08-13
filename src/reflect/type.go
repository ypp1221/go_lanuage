package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Sex  string
	Age  int
}

func (p person) GetMethod() {
	fmt.Println("method")
}
func (p person) PrivateMethod() {

}
func GetTypeInformation(a interface{}) {
	t := reflect.TypeOf(a)

	fmt.Println("a的类型信息是：", t)
	k := t.Kind()

	fmt.Println("a的底层类型是 :", k)
	switch k {
	case reflect.Float64:
		fmt.Println("float64")
	case reflect.String:
		fmt.Println(" string")
	}
}

func GetInterfanceType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("a的类型信息为:", t)
	k := t.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println("底层类型是:float64")
	case reflect.String:
		fmt.Println("底层类型是:string")
	}
}
func GetValue(a interface{}) {

	t := reflect.ValueOf(a)
	k := t.Kind()
	switch k {
	case reflect.Float64:
		t.SetFloat(9.9)
	case reflect.String:
		t.SetString("yang")
	case reflect.Ptr:
		t.Elem().SetFloat(8.9)
	}
	fmt.Println(t.Elem().Float())
}

func GetPersonInf(a interface{}) {
	ty := reflect.TypeOf(a)
	fmt.Println("ty类型:", ty)
	fmt.Println("stirng类型:", ty.Name())
	va := reflect.ValueOf(a)
	fmt.Println(va)
	for i := 0; i < ty.NumField(); i++ {
		temp := ty.Field(i)
		fmt.Printf("name:%s, type:%v,", temp.Name, temp.Type)

		val := va.Field(i).Interface()
		if val == "yang" {
			va.Field(i).SetString("zhang")
		}
		fmt.Printf("value:%v\n", val)
	}

	for i := 0; i < ty.NumMethod(); i++ {
		m := ty.Method(i)
		fmt.Println("name:", m.Name, "type:", m.Type)
	}

}
func main() {
	//var a float64 = 3
	//GetTypeInformation(a)
	//var str string = "what"
	//GetInterfanceType(a)
	//GetInterfanceType(str)
	//值进行修改时，必须时指针类型
	//GetValue(&a)
	a := person{
		"yang",
		"men",
		22,
	}
	GetPersonInf(a)
}
