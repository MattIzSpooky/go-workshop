## 3. Goroutines
Go routines are one of the things that makes Go special.
They are lightweight threads managed by the Go runtime. 

A Go routine can be started by prefixing a function with the ```go``` keyword.

A regular function call would be:
```func()```

The same call but ran in a Go routine would be:
```go func()```

Careful though, ***do not start a go routine without knowing when or if it will end. This will result in memory leak-like problems.***

[This article has more information on this](https://dave.cheney.net/2016/12/22/never-start-a-goroutine-without-knowing-how-it-will-stop)

### Communicating between Go routines
Communicating between go routines is done through Channels. 
Imagine a channel as a pipe you can send data over, it's that simple.

See the code for more information.

Channels can be quite deep. There are buffered and unbuffered channels.
You can also define whether a function can only read or write to a channel

See [main.go](main.go) for a full example.