package main

type ak47 struct {
	gun
}

func newak47() iGun {
	return &ak47{
		gun: gun{
			"AK47 gun",
			4,
		},
	}
}
