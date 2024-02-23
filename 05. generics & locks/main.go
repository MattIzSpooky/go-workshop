package main

import (
	"fmt"
	"workshop/generics+locks/algorithms"
	"workshop/generics+locks/person"
	"workshop/generics+locks/structures"
)

func main() {

	fmt.Println("Queue example")
	fmt.Println()
	queue := structures.Queue[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for !queue.IsEmpty() {
		queue.Dequeue()
	}

	fmt.Println()
	fmt.Println("Stack example")
	fmt.Println()

	stack := structures.Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for !stack.IsEmpty() {
		stack.Pop()
	}

	fmt.Println()
	fmt.Println("Linked list example")
	fmt.Println()

	linkedList := structures.LinkedList[*person.Person]{}
	linkedList.Insert(person.New("Matthijs"))
	linkedList.Insert(person.New("Joseph"))
	linkedList.Insert(person.New("Jake"))
	linkedList.Insert(person.New("Vincent"))
	linkedList.Insert(person.New("Joe"))

	linkedList.Print()

	resultFound := linkedList.FindFunc(func(p *person.Person) bool {
		return p.Name == "Jake"
	})

	resultNotFound := linkedList.FindFunc(func(p *person.Person) bool {
		return p.Name == "Alex"
	})

	fmt.Println(fmt.Sprintf("FindFunc resultFound: %v", resultFound))
	fmt.Println(fmt.Sprintf("FindFunc resultNotFound: %v", resultNotFound))

	fmt.Println()
	fmt.Println("Binary search & Quick sort")
	fmt.Println()

	myList := algorithms.List[person.Person]{}

	myList.Add(person.NewWithId(1, "Matthijs"))
	myList.Add(person.NewWithId(2032, "Jerome"))
	myList.Add(person.NewWithId(23, "Alex"))
	myList.Add(person.NewWithId(93, "Joseph"))
	myList.Add(person.NewWithId(10203, "Jake"))
	myList.Add(person.NewWithId(99, "Blake"))
	myList.Add(person.NewWithId(6372, "Rosa"))
	myList.Add(person.NewWithId(938183, "Vera"))
	myList.Add(person.NewWithId(2923, "Vince"))
	myList.Add(person.NewWithId(211, "Ashley"))

	myList.Print()

	fmt.Println("Everything was inserted in order!")
	fmt.Println("And now we randomize!")

	myList.Randomize()
	myList.Print()

	fmt.Println("And now we sort!")

	myList.Sort()
	myList.Print()

	fmt.Println("Now let us do a search: ")
	result := myList.Search(10203)

	fmt.Println(fmt.Sprintf("Search result: [%v]", *result))

}
