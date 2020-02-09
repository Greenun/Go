package main

import (
	"fmt"
)

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue{
	return &Queue{elements: make([]interface{}, 0)}
}

func (queue *Queue) Enqueue(data interface{}){
	queue.elements = append(queue.elements, data)
}

func (queue *Queue) Dequeue() interface{}{
	if len(queue.elements) < 1 {
		panic("Queue is Empty")
	}
	var value interface{}
	value, queue.elements = queue.elements[:1], queue.elements[1:]
	return value
}

func (queue *Queue) Length() int {
	return len(queue.elements)
}

func main(){
	var q *Queue = NewQueue()
	q.Enqueue("test")
	q.Enqueue(10)
	fmt.Println(q.Length())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
