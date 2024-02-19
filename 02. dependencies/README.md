## 2. Dependencies

One of the most popular things in modern software development is the usage of dependencies.
There is little to no software out there that uses no external dependencies.

Go can also make use of external dependencies. 
Though it is heavily simplified when compared to dependency management tools such as Maven, Gradle, npm or NuGet.
Its simplicity makes and breaks dependency management in Go. You can simply download code from a repository and start coding right away.

To start keeping track of dependencies one must run the following command to start using go modules.

```
go mod init [whatever]
```

This will create a `go.mod` file that keeps track of all dependencies in the project.

it should look something like

```
module workshop/dependencies

go 1.21.6
```

Now we will add a dependency. We will add https://pkg.go.dev/github.com/fatih/color#section-readme.

This can be done by running 
```go get github.com/fatih/color```

We can now start using the color library in our program.

Note how "fatih/color" is listed as an indirect package. 
After using it our code and running ```go mod tidy``` it will be listed as a primary dependency.

Want to know more about dependency management in Go? Read more on https://go.dev/doc/modules/managing-dependencies