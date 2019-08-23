package main

import "errors"

type Queue struct {
	maxSize int
	array   [4]int
	front   int
	rear    int
}

func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) ShowQueue() {

}

func main() {

}
