# Any | Dict | List

## Python Functions

```python
def method_args(*args): pass
def method_kwargs(**kwargs): pass
def method_pythonic(*args, **kwargs): pass
```

## Go Functions

```go
package main

import "fmt"

// Any is a generic type `alias` for any value.
type Any = interface{}

// Dict is a type `alias` for a map with string keys and values of any type.
type Dict = map[string]interface{}

// List is a type `alias` for a slice of values of any type.
type List = []interface{}

// methodArgs simulates Python's *args by accepting a variadic list of arguments.
func methodArgs(args ...Any) {
 fmt.Println("Handling *args:")
 for i, arg := range args {
  fmt.Printf("  Arg %d: %v\n", i, arg)
 }
}

// methodKwargs simulates Python's **kwargs by accepting a dictionary of key-value pairs.
func methodKwargs(kwargs Dict) {
 fmt.Println("Handling **kwargs:")
 for key, value := range kwargs {
  fmt.Printf("  %s: %v\n", key, value)
 }
}

// methodPythonic combines *args and **kwargs functionality, accepting both positional and keyword arguments.
func methodPythonic(args List, kwargs Dict) {
 // Handle *args
 fmt.Println("Handling *args:")
 for i, arg := range args {
  fmt.Printf("  Arg %d: %v\n", i, arg)
 }

 // Handle **kwargs
 fmt.Println("Handling **kwargs:")
 for key, value := range kwargs {
  fmt.Printf("  %s: %v\n", key, value)
 }
}

func main() {
 // Example usage of methodArgs
 methodArgs(1, "hello", true, 42.5)

 // Example usage of methodKwargs
 methodKwargs(Dict{
  "name":  "Alice",
  "age":   30,
  "admin": true,
 })

 // Example usage of methodPythonic
 methodPythonic(
  List{10, "example", false},
  Dict{
   "status":  "active",
   "message": "success",
  },
 )
}
```
