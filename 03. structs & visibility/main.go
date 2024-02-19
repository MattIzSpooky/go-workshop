package main

import (
	"fmt"
	"workshop/structs+visibility/pets"
)

func main() {
	scooby := pets.NewDog("Scooby")
	jake := pets.NewCat("jake")

	scooby.MakeNoise()

	errWalkScooby := scooby.Walk()
	if errWalkScooby != nil {
		scooby.GiveCookie()

		// We choose to ignore the error this time.
		_ = scooby.Walk()
	}

	fmt.Println()

	errWalkJake := jake.Walk()
	if errWalkJake != nil {
		fmt.Println("Right... Cats are lazy")
	}

	jake.MakeNoise()

	clonedCat := jake.Clone()
	fmt.Println(fmt.Sprintf("Cloned cat!! %v", clonedCat))
}
