package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// This is how you declare a const in Go.
const sleepTimer time.Duration = 50
const timesToPrint int = 5

func printXAmountOfTimes(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(sleepTimer * time.Millisecond)
		fmt.Println(s)
	}
}

func generate10RandomNumbers(id int, c chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	for i := 0; i < 10; i++ {
		// Put a string into the channel
		c <- fmt.Sprintf("Worker %d: %d", id, rand.Int())
	}

	fmt.Printf("Worker %d done\n", id)
	wg.Done() // Declare a task to be finished, this will remove one from the waiting list.
	// It essentially does wg.Add(-1)
}

func main() {
	go printXAmountOfTimes("i print asynchronous", timesToPrint)
	printXAmountOfTimes("i print synchronously", timesToPrint)

	fmt.Println("=======================")

	c := make(chan string)    // This creates an unbuffered channel.
	wg := new(sync.WaitGroup) // This instantiates a wait group
	// WaitGroups are useful for running a bunch of tasks and waiting till all are done.

	// Start a routine to read from the channel
	go func() {
		// In theory, if this channel never closes, the go routine will cause a memory leak.
		// When the Go runtime exits it will close all channels, however, one cannot rely upon this for long-running processes such as HTTP servers.
		// Therefore, it's best to close channels explicitly.

		// Don't worry, this go routine will exit when the channel gets closed later in the program.

		// Here we will read from each value coming from the channel
		for item := range c {
			fmt.Println(item)
		}
	}()

	// Run 5 go routines to generate a bunch of random numbers
	for i := 0; i < 5; i++ {
		wg.Add(1) // This adds the amount of tasks that need to be done.

		// Start the asynchronous task
		go generate10RandomNumbers(i, c, wg)
	}

	// This will block till all tasks are finished.
	wg.Wait()

	// Close the channel explicitly.
	close(c)

	fmt.Println("=======================")

	// Below is a little trick that can be used to keep an application alive, such as an HTTP-server.

	sigChan := make(chan os.Signal, 1) // Create a buffered channel that will automatically close once 1 element gets put in it.

	// We will listen to some specific OS commands sent to the application
	// The sigChan will be filled when one of those commands is sent to the application.
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	color.Green("The application will continue to run until a stop signal (CTRL+C in terminal) is sent to the application.")

	<-sigChan // Reading from a channel is blocking! This can be really useful for communicating between go routines.
}
