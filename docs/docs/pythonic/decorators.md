# Decorators

Go doesn't have decorators or wrappers in the same way Python does, but you can achieve similar functionality using **higher-order functions** (functions that take other functions as arguments or return them), **closures**, and **interfaces**.

## Examples

### Achieving Decorator-like Behavior in Go

1. **Higher-order functions**: In Go, you can pass functions as arguments to other functions, or return them. This lets you wrap or modify the behavior of a function.

2. **Closures**: You can create closures (functions that "remember" the environment in which they were created) that allow for wrapping functionality and preserving state.

### Example 1: Wrapping Functions (like a decorator)

```go
package main

import "fmt"

// A function that takes another function and adds additional behavior (like a decorator)
func logExecution(fn func(string) string) func(string) string {
    return func(input string) string {
        fmt.Println("Before calling the function")
        result := fn(input)  // Call the original function
        fmt.Println("After calling the function")
        return result
    }
}

// A simple function to be decorated
func sayHello(name string) string {
    return "Hello, " + name
}

func main() {
    // Wrap sayHello with the logExecution "decorator"
    decorated := logExecution(sayHello)

    // Call the decorated function
    result := decorated("Alice")
    fmt.Println(result)  // Outputs: "Hello, Alice"
}
```

### Explanation 1

- `logExecution` is a higher-order function that takes another function `fn` as an argument and returns a new function that wraps `fn`. This is like a decorator in Python.
- The returned function has additional logic (in this case, logging before and after the function call) around the original function.

### Example 2: Wrapping Methods with Closures

In Go, closures can also be used to encapsulate behavior around methods.

```go
package main

import "fmt"

// A function that takes a method and wraps it with additional behavior
func addLogging(fn func()) func() {
    return func() {
        fmt.Println("Before function call")
        fn()  // Original function call
        fmt.Println("After function call")
    }
}

func greet() {
    fmt.Println("Hello, World!")
}

func main() {
    // Wrap greet with addLogging
    loggedGreet := addLogging(greet)
    loggedGreet()  // Outputs logs before and after greet
}
```

### Example 3: Wrapping Methods with Interfaces

You can also achieve decorator-like functionality with **interfaces** in Go, especially when working with methods in structs.

```go
package main

import "fmt"

// Define an interface for the behavior we want to decorate
type Greeter interface {
    Greet() string
}

// A basic struct that implements the Greeter interface
type Person struct {
    Name string
}

func (p Person) Greet() string {
    return "Hello, " + p.Name
}

// A decorator that wraps a Greeter and modifies its behavior
type GreeterWithLog struct {
    Greeter
}

func (g GreeterWithLog) Greet() string {
    fmt.Println("Before greeting")
    result := g.Greeter.Greet()  // Call the original Greet method
    fmt.Println("After greeting")
    return result
}

func main() {
    person := Person{Name: "Alice"}
    greeter := GreeterWithLog{Greeter: person}

    fmt.Println(greeter.Greet())  // Outputs log and the greeting
}
```

### Explanation 3

- **Interface**: We define the `Greeter` interface and then implement it with the `Person` struct.
- **Decorator**: The `GreeterWithLog` struct "wraps" a `Greeter` and modifies its behavior. The method `Greet()` of `GreeterWithLog` adds additional behavior before and after calling the original `Greet()` method.

### Differences from Python Decorators

- **Python decorators**: Python decorators are syntactic sugar that allows wrapping a function or method in a cleaner way using the `@decorator` syntax.
- **Go Wrappers**: In Go, decorators can be implemented using higher-order functions, closures, or interfaces, but the syntax and approach differ from Python.

### Summary

While Go doesn't have Python-style decorators, you can still achieve similar functionality through higher-order functions, closures, and interfaces, which are powerful tools for wrapping and modifying behavior in Go.
