package main

import (
	"fmt"
	"math/rand"
	"workshop/generics+locks/algorithms"
	"workshop/generics+locks/structures"
)

type Person struct {
	id   int
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

	fmt.Println()

	fmt.Println("Binary search & Quick sort")
	myList := algorithms.List[Person]{}

	myList.Add(createPersonWithId(23, "Matthijs"))
	myList.Add(createPersonWithId(1, "Alex"))
	myList.Add(createPersonWithId(93, "Joseph"))
	myList.Add(createPersonWithId(10203, "Jake"))

	myList.Print()
	fmt.Println("And now we sort!")
	myList.Sort()
	myList.Print()

	result := myList.Search(23)

	fmt.Println(fmt.Sprintf("Search result: [%v]", *result))

}

func (p Person) ID() int {
	return p.id
}

func createPerson(name string) *Person {
	return &Person{
		id:   rand.Int(),
		Name: name,
	}
}

func createPersonWithId(id int, name string) *Person {
	return &Person{
		id:   id,
		Name: name,
	}
}
