package main

type nike struct {
}

func (n *nike) makeShop() iShop {
	return &nikeShop{
		Shop: Shop{
			"nike",
			23,
		},
	}
}
func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			"nike",
			23,
		},
	}
}
