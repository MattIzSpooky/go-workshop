package main

import (
	"fmt"
	"math/rand"
	"workshop/generics+locks/structures"
)

type Person struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Queue example")

	queue := structures.Queue[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for !queue.IsEmpty() {
		queue.Dequeue()
	}

	fmt.Println("Stack example")

	stack := structures.Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for !stack.IsEmpty() {
		stack.Pop()
	}

	fmt.Println("Linked list example")
	linkedList := structures.LinkedList[*Person]{}
	linkedList.Insert(createPerson("Matthijs"))
	linkedList.Insert(createPerson("Joseph"))
	linkedList.Insert(createPerson("Jake"))
	linkedList.Insert(createPerson("Vincent"))
	linkedList.Insert(createPerson("Joe"))

	linkedList.Print()

	resultFound := linkedList.FindFunc(func(p *Person) bool {
		return p.Name == "Jake"
	})

	resultNotFound := linkedList.FindFunc(func(p *Person) bool {
		return p.Name == "Alex"
	})

	fmt.Println(fmt.Sprintf("FindFunc resultFound: %v", resultFound))
	fmt.Println(fmt.Sprintf("FindFunc resultNotFound: %v", resultNotFound))
}

func createPerson(name string) *Person {
	return &Person{
		ID:   rand.Int(),
		Name: name,
	}
}
