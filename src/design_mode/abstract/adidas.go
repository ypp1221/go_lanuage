package main

type adidas struct {
}

func (a *adidas) makeShop() iShop {
	return &adidasShop{
		Shop: Shop{
			"adidas",
			14,
		},
	}
}
func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			"adidas",
			23,
		},
	}
}
