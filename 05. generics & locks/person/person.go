package person

import "math/rand"

type Person struct {
	id   int
	Name string
}

func (p Person) ID() int {
	return p.id
}

func New(name string) *Person {
	return &Person{
		id:   rand.Int(),
		Name: name,
	}
}

func NewWithId(id int, name string) *Person {
	return &Person{
		id:   id,
		Name: name,
	}
}
