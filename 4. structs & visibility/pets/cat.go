package pets

import (
	"errors"
	"fmt"
)

type Cat struct {
	Pet                // This is called type embedding. "Cat" will inherit all properties and methods from "Pet"
	wantsToBeLeftAlone bool
}

func NewCat(name string) *Cat {
	return &Cat{
		Pet:                Pet{name: name},
		wantsToBeLeftAlone: false}
}

// Interfaces are implicitly implemented
func (c *Cat) MakeNoise() {
	if c.wantsToBeLeftAlone {
		fmt.Printf("Cat [%s] hisses!", c.name)
	} else {
		fmt.Printf("Cat [%s] says meow!", c.name)
	}

	c.wantsToBeLeftAlone = generateBool()
}

func (c *Cat) Walk() error {
	c.wantsToBeLeftAlone = generateBool()

	return errors.New(fmt.Sprintf("Cat with name [%s] prefers to sleep instead...", c.name))
}
