package main

import (
	"fmt"
	"sync"
)

type linknode struct {
	value int
	next  *linknode
}
type link_stack struct {
	root *linknode
	size int
	lock sync.Mutex
}

func (stack *link_stack) Push(element int) {

	stack.lock.Lock()
	defer stack.lock.Unlock()
	NewRoot := new(linknode)
	NewRoot.value = element
	if stack == nil {
		//空栈,创建
		stack.root = NewRoot
	} else {
		PreNode := stack.root
		NewRoot.next = PreNode
		stack.root = NewRoot
	}
	stack.size++
}
func (stack *link_stack) Empty() int {
	if stack.size == 0 {
		return 1
	}
	return 0
}
func (stack *link_stack) Pop() (element *linknode) {
	stack.lock.Lock()
	stack.lock.Unlock()
	if stack.Empty() == 1 {
		panic("栈是空")
	}
	element = stack.root
	stack.root = stack.root.next
	return element
}
func (stack *link_stack) Top() (element *linknode) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	element = stack.root
	return element
}

func main() {
	sta := new(link_stack)
	sta.Push(2)
	sta.Push(4)
	sta.Push(9)
	sta.Push(8)
	fmt.Println(sta.Top().value)
	sta.Pop()
	fmt.Println(sta.Top().value)

}
