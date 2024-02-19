## 3. Structs & Visibility

Go does not have classes, but with structs you can define your types.
In the very basics, a struct is a collection of fields.
One may add methods to a struct. In Go, these are called functions or "receiver" functions.

```go
package main

import (
	"fmt"
	"math"
)

// This is how you define a struct in Go.
type MyVector struct {
	X float64
	Y float64
}

// This is how you add a function to a struct
// The "(v Vector)" indicates the "receiver". It is kind of like the "this" keyword in Java or C#
func (v MyVector) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// This is how you instantiate an object
    v := MyVector{X: 3.0, Y: 4.0}
    fmt.Println(v.Abs())
}
```

### Visibility
Are there keywords such as "private", "public" in Go?
No, but the functionality is there. Structs and functions that start with a capital letter (UpperCamelCase/PascalCase), 
such as in the "MyVector" example are public outside the current package.
Methods that start with a lower letter (camelCase) are private and can only be used within the current package.
This applies for fields too, but beware. 
Private/public only applies to packages, meaning that you can modify the private values or call private methods of other objects within the same package.
If you want to modify an object from outside its package you will have to add getters and setters to the struct.

### Embedding
Go does not have classes and you cannot extend other objects.
However, you can "embed" structs within other structs. 
With this you can copy functionality of a struct into the other struct.
You can embed as many structs as you like.

One thing to keep in mind is that embedding is not the same as inheritance.
You cannot do "instanceof" checks to check if an object is of a certain type.

Embedding (composition) adds a lot of benefits over inheritance.
I am not going to explain why but wikipedia has a decent page on it https://en.wikipedia.org/wiki/Composition_over_inheritance

Read more here
https://go.dev/doc/effective_go#embedding

This example contains both visibility, embedding and interfaces.

### Pointers
Go can use pointers, however, be careful when using them. 
Due to Go's nature to easily spin up threads it is preferred to copy values instead of referring to them.
However, if your struct contains something that is not cheap or not viable to copy,
such as a buffer, an SQL connection, a mutex, etc. use a pointer.

Another reason to use a pointer is when you want to modify an object.

You can read more here: https://go.dev/doc/effective_go#pointers_vs_values

Luckily, you do not have to clean up pointers yourself. Once all references to an object are lost, 
Go will clean it right up.
