package main

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			"Musket gun",
			8,
		},
	}
}
