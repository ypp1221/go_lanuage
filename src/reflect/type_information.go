package reflect

import (
	"fmt"
	"reflect"
)

func GetTypeInformation(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("a的类型信息是：", t)
	k := t.Kind()

	fmt.Println("k is :", k)
	switch k {
	case reflect.Float64:
		fmt.Println("k's type is float64")
	case reflect.String:
		fmt.Println("k's type i string")
	}
}
func main() {
	var a float64 = 3
	GetTypeInformation(a)
}
