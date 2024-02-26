## 5. Generics & Locks

Like Java and C#, Golang supports generics. What is nice about generics in Golang is that like with generics 
in C# the generic type information is not lost during compile time.

The devs wrote a great introductory article about generics.
https://go.dev/blog/intro-generics

This part contains a bunch of examples where generics are used.
See [main.go](main.go) for a full example.

### Locks
Locks are incredibly simple... in theory.
Locks should be introduced when you share state between multiple go routines.
This protects the state of variable and helps with race conditions.
Whenever you lock a lock (heh), the next call requesting the lock will have to wait until the locks is freed.
This means that your go routine will block until the lock is freed. Not freeing the lock can cause your application to wait
for the lock to come free indefinitely. So, be careful with locks.

Locks have been used in this example but have been better used in the [gRPC example](../11.%20gRPC).