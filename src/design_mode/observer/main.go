package main

func main() {
	shritItem := newItem("shrit")
	observerFirst := &customer{id: "yang@qq.com"}
	observerSecond := &customer{id: "zhang@163.com"}
	shritItem.register(observerFirst)
	shritItem.register(observerSecond)
	shritItem.updateAvailability()
}
