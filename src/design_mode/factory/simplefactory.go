package main

type iGun interface {
	setName(name string)
	getName() string
	setPower(power int)
	getPower() int
}
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}
func (g *gun) getName() string {
	return g.name
}
func (g *gun) setPower(power int) {
	g.power = power
}
func (g *gun) getPower() int {
	return g.power
}
