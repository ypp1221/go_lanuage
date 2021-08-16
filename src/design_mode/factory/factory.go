package main

func getGun(guntype string) iGun {
	if guntype == "ak47" {
		return newak47()
	} else if guntype == "musket" {
		return newMusket()
	}
	return nil
}
