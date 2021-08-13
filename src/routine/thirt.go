package main

/*
int Add(int a, int b){
	return a + b;
}
*/
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println("测试是否提交成功")
	//pack.Test(0)
	fmt.Println(C.Add(1, 2))
}
