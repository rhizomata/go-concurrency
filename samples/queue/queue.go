package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// Queue ...
type Queue struct {
	sync.Mutex
	innerList *list.List
	lock      chan bool
	waiting   int64
}

//NewQueue ...
func NewQueue() *Queue {
	queue := Queue{innerList: list.New(), lock: make(chan bool)}
	return &queue
}

// Push ..
func (queue *Queue) Push(value interface{}) {
	queue.Lock()
	queue.innerList.PushBack(value)
	if queue.waiting > 1 {
		queue.lock <- false
	}
	queue.Unlock()
}

// Pop ..
func (queue *Queue) Pop() (value interface{}) {
	queue.waiting = queue.waiting + 1
	if queue.innerList.Len() == 0 {
		<-queue.lock
	}
	queue.Lock()
	queue.waiting = queue.waiting - 1
	el := queue.innerList.Front()
	value = el.Value
	queue.innerList.Remove(el)

	queue.Unlock()
	return value
}

func main() {
	queue := NewQueue()
	quit := make(chan bool)

	for i := 0; i < 100; i++ {
		queue.Push(fmt.Sprintf("Default %d", i))
	}

	fmt.Println("Pre 1", queue.Pop())
	fmt.Println("Pre 2", queue.Pop())
	// quit <- true

	go func(queue *Queue) {
		// <-quit
		for i := 0; i < 100; i++ {
			queue.Push(fmt.Sprintf("Send %d", i))
			time.Sleep(2 * time.Millisecond)
			if i%10 == 5 {
				time.Sleep(200 * time.Millisecond)
			}
		}
		quit <- true
	}(queue)

	fmt.Println("main 1", queue.Pop())

	go func(queue *Queue) {
		for true {
			fmt.Println("Rec:1 > ", queue.Pop())
		}
	}(queue)

	fmt.Println("main 2", queue.Pop())
	go func(queue *Queue) {
		for true {
			fmt.Println("Rec:2 > ", queue.Pop())
		}
	}(queue)

	go func(queue *Queue) {
		for true {
			fmt.Println("Rec:3 > ", queue.Pop())
		}
	}(queue)

	<-quit
}
