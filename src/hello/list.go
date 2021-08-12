package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}
type Double_Link struct {
	data      int
	next, pre *Double_Link
}

//创建一个单链表
func create_link(length int) *LinkNode {
	first := new(LinkNode)
	first.data = -1
	cur := first
	for i := 0; i < length; i++ {
		node := new(LinkNode)
		node.data = i

		node.next = cur.next
		cur.next = node
		cur = node
	}
	return first
}

//获取单链表的长度
func (first *LinkNode) get_link_length() (length int) {
	length = 0
	for first != nil {
		length++
		first = first.next
	}
	return length
}
func print_link(first *LinkNode) {
	for first != nil {
		fmt.Printf("%v ", first.data)
		first = first.next
	}
	fmt.Println()
}

func create_double_link(length int) *Double_Link {
	//创建链表头结点
	first := new(Double_Link)
	first.data = -1
	first.next = first
	first.pre = first

	cur := first
	for i := 0; i < 10; i++ {
		node := new(Double_Link)
		node.data = i
		node.next = cur.next
		cur.next = node
		node.pre = cur.pre
		cur.pre = node
		cur = node

	}
	return first
}
func (first *Double_Link) print_doulbe() {
	cur := first
	for cur.next != first {
		fmt.Printf("%v ", cur.data)
		cur = cur.next
	}
	fmt.Printf("%v ", cur.data)
}
func main() {
	first := create_link(10)
	print_link(first)
	fmt.Println("single linknode length is :", first.get_link_length())
	double_first := create_double_link(10)
	double_first.print_doulbe()
	//print_doulbe(double_first)
}
