package main

import "fmt"

func printShopDetails(s iShop) {
	fmt.Println(s.getLogo(), s.getSize())
}
func main() {
	adidasFactory, _ := getSportFactory("adidas")
	nikeFactory, _ := getSportFactory("nike")

	adidasShop := adidasFactory.makeShop()
	nikeShop := nikeFactory.makeShirt()
	printShopDetails(adidasShop)
	printShopDetails(nikeShop)

}
