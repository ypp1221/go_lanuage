package main

import (
	"fmt"
)

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}
func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}
func (i *item) deregister(o observer) {
	i.observerList = removeFromslice(i.observerList, o)
}
func (i *item) notifyAll() {
	for _, obser := range i.observerList {
		obser.update(i.name)
	}
}
func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerLength := len(observerList)
	for i, obser := range observerList {
		if obser == observerToRemove {
			observerList[i-1], observerList[observerLength-1] = observerList[observerLength-1], observerList[i-1]
		}
	}
	return observerList
}
