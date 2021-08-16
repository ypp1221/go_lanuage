package main

type subject interface {
	register(Observer observer)
	deregister(observer observer)
	notifyAll()
}
