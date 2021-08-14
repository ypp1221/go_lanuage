package main

/*
int Add(int a, int b){
	return a + b;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.Add(1, 2))

}
