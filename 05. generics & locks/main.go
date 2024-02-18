package main

import (
	"fmt"
	"workshop/generics+locks/structures"
)

func main() {
	fmt.Println("Queue example")

	queue := structures.Queue[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for !queue.IsEmpty() {
		item := queue.Dequeue()
		fmt.Println(item)
	}

	fmt.Println("Stack example")

	stack := structures.Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Println(item)
	}
}
