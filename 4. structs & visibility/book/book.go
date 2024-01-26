package book

import (
	"fmt"
	"time"
)

// This is our struct that contains data
type Book struct {
	name   string
	author string
	pages  int
	genres []string
}

const sleepDuration time.Duration = 100

// Here we will list our struct methods.
func (b Book) Read() {
	fmt.Printf("I have started reading my book.")

	time.Sleep(sleepDuration * time.Millisecond)
	for i := 0; i < b.pages; i++ {
		fmt.Print("zzz..")

		time.Sleep(time.Duration(i) * sleepDuration * time.Second)
	}
	defer func() {
		fmt.Println("I am done reading ")
	}()
}
