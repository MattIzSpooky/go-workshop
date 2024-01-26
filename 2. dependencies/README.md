## 2. Dependencies

To start keeping track of dependencies one must run the following command to start using go modules.

```
go mod init [whatever]
```

This will create a `go.mod` file that keeps track of all dependencies.

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