package pets

import (
	"errors"
	"fmt"
)

type Cat struct {
	Pet                     // This is called type embedding. "Cat" will inherit all properties and methods from "Pet"
	wantsToBeLeftAlone bool // This is a private value of Cat
}

// Private values cannot be set from the outside. A common pattern is exposing a "New" function.
func NewCat(name string) *Cat {
	return &Cat{
		Pet:                Pet{name: name},
		wantsToBeLeftAlone: false}
}

// Interfaces are implicitly implemented. The NoiseMaker interface is implemented here.
func (c *Cat) MakeNoise() {
	if c.wantsToBeLeftAlone {
		fmt.Println(fmt.Sprintf("Cat [%s] hisses!", c.name))
	} else {
		fmt.Println(fmt.Sprintf("Cat [%s] says meow!", c.name))
	}

	c.wantsToBeLeftAlone = generateBool()
}

// Interfaces are implicitly implemented. The Walk interface is implemented here.
func (c *Cat) Walk() error {
	c.wantsToBeLeftAlone = generateBool()

	return errors.New(fmt.Sprintf("Cat with name [%s] prefers to sleep instead...", c.name))
}

func (c *Cat) Clone() *Cat {
	return &Cat{
		Pet:                Pet{name: c.name},
		wantsToBeLeftAlone: c.wantsToBeLeftAlone,
	}
}
