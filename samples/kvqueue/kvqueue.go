package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// KVQueue ...
type KVQueue struct {
	sync.Mutex
	innerList *list.List
	kvMap     map[interface{}]interface{}
	lock      chan bool
	waiting   int64
}

//NewKVQueue ...
func NewKVQueue() *KVQueue {
	queue := KVQueue{innerList: list.New(), kvMap: make(map[interface{}]interface{}), lock: make(chan bool)}
	return &queue
}

// Push ..
func (queue *KVQueue) Push(key interface{}, value interface{}) {
	queue.Lock()
	oldval := queue.kvMap[key]
	if oldval == nil {
		queue.innerList.PushBack(key)
	} else {
		fmt.Println("----------------- over write : ", key, oldval, value)
	}
	queue.kvMap[key] = value

	if queue.waiting > 1 {
		queue.lock <- false
	}
	queue.Unlock()
}

// Pop ..
func (queue *KVQueue) Pop() (key interface{}, value interface{}) {
	key = queue._pop()

	for ; key == nil; key = queue._pop() {
		queue.waiting = queue.waiting + 1
		<-queue.lock
		queue.waiting = queue.waiting - 1
	}

	value = queue.kvMap[key]
	return key, value
}

// Pop ..
func (queue *KVQueue) _pop() (key interface{}) {
	queue.Lock()
	el := queue.innerList.Front()
	if el != nil {
		key = el.Value
		queue.innerList.Remove(el)
	}
	queue.Unlock()
	return key
}

func main() {
	queue := NewKVQueue()
	// quit := make(chan bool)

	for i := 0; i < 100; i++ {
		queue.Push(fmt.Sprintf("A%d", i), fmt.Sprintf("Pre %d", i))
	}

	k, v := queue.Pop()
	fmt.Println("Pre 1", k, v)
	k, v = queue.Pop()
	fmt.Println("Pre 2", k, v)
	// quit <- true

	go func(queue *KVQueue) {
		// <-quit
		for i := 0; i < 100; i++ {
			queue.Push(fmt.Sprintf("B%d", i), fmt.Sprintf("Send %d", i))
			time.Sleep(2 * time.Millisecond)
			if i%10 == 5 {
				time.Sleep(10 * time.Millisecond)
			}
		}
		// quit <- true
	}(queue)

	k, v = queue.Pop()
	fmt.Println("main 1", k, v)

	go func(queue *KVQueue) {
		for true {
			k, v := queue.Pop()
			fmt.Println("Rec:1 > ", k, v)
		}
	}(queue)

	go func(queue *KVQueue) {
		for i := 0; i < 100; i++ {
			queue.Push(fmt.Sprintf("B%d", i), fmt.Sprintf("Mid %d", i))
			time.Sleep(2 * time.Millisecond)
		}
	}(queue)

	k, v = queue.Pop()
	fmt.Println("main 2", k, v)
	go func(queue *KVQueue) {
		for true {
			k, v := queue.Pop()
			fmt.Println("Rec:2 > ", k, v)
		}
	}(queue)

	go func(queue *KVQueue) {
		for true {
			k, v := queue.Pop()
			fmt.Println("Rec:3 > ", k, v)
		}
	}(queue)

	// <-quit
	for i := 0; i < 500; i++ {
		queue.Push(fmt.Sprintf("A%d", i), fmt.Sprintf("Post %d", i))
		time.Sleep(2 * time.Millisecond)
	}

	// time.Sleep(10 * time.Second)
}
