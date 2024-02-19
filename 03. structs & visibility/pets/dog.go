package pets

import (
	"errors"
	"fmt"
)

type Dog struct {
	Pet             // This is called type embedding. "Dog" will inherit all properties and methods from "Pet"
	requiresACookie bool
}

// Private values cannot be set from the outside. A common pattern is exposing a "New" function.
func NewDog(name string) *Dog {
	return &Dog{
		Pet:             Pet{name: name},
		requiresACookie: false}
}

// Interfaces are implicitly implemented. The NoiseMaker interface is implemented here.
func (d *Dog) MakeNoise() {
	d.requiresACookie = generateBool()

	if d.requiresACookie {
		fmt.Println(fmt.Sprintf("Dog [%s] requires a cookie. It isn't gonna do much without...", d.GetName()))
	} else {
		fmt.Println(fmt.Sprintf("Dog [%s] says woof!", d.GetName()))
	}
}

// Interfaces are implicitly implemented. The Walk interface is implemented here.
func (d *Dog) Walk() error {
	if d.requiresACookie {
		return errors.New(fmt.Sprintf("Dog with name [%s] does not want to walk unless they get a cookie", d.GetName()))
	}

	fmt.Println(fmt.Sprintf("Dog [%s] is happy to walk..", d.GetName()))

	return nil // Since the function expects we return an error, we can return nil instead. signaling that there was no error executing this function.
}

func (d *Dog) GiveCookie() {
	d.requiresACookie = false

	fmt.Println(fmt.Sprintf("The monster [%s] has been fed...", d.GetName()))
}

func (c *Dog) Clone() *Dog {
	return &Dog{
		Pet:             Pet{name: c.name},
		requiresACookie: c.requiresACookie,
	}
}
