package main

//抽象产品
type iShop interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

//具体产品
type Shop struct {
	logo string
	size int
}

//实现抽象产品的方法
func (s *Shop) setLogo(logo string) {
	s.logo = logo
}
func (s *Shop) setSize(size int) {
	s.size = size
}
func (s *Shop) getLogo() string {
	return s.logo
}
func (s *Shop) getSize() int {
	return s.size
}
