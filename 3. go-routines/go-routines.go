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

const sleepTimer time.Duration = 50
const timesToPrint int = 5

func printXAmountOfTimes(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(sleepTimer * time.Millisecond)
		fmt.Println(s)
	}
}

func generate10RandomNumbers(id int, c chan string, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("Worker %d: %d", id, rand.Int())
	}

	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func main() {
	go printXAmountOfTimes("i print asynchronous", timesToPrint)
	printXAmountOfTimes("i print synchronously", timesToPrint)

	fmt.Println("=======================")

	c := make(chan string)
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go generate10RandomNumbers(i, c, wg)
	}

	go func() {
		// In theory, if this channel never closes, the go routine will cause a memory leak.
		// When the Go runtime exits it will close all channels, however, one cannot rely upon this for long-running processes such as HTTP servers.
		// Therefore, it's best to close channels explicitly.
		for item := range c {
			fmt.Println(item)
		}
	}()

	wg.Wait()

	close(c)

	fmt.Println("=======================")

	// Below is a little trick that can be used to keep an application alive, such as an HTTP-server.

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGHUP,
	)

	color.Green("The application will continue to run until a stop signal is sent to the application.")

	<-sigChan // Reading from a channel is blocking! This can be really useful for communicating between go routines.
}
