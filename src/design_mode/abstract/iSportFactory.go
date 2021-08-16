package main

import "fmt"

type iSportFactory interface {
	makeShop() iShop
	makeShirt() iShirt
}

func getSportFactory(brand string) (iSportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("没有对应的类型")
}
