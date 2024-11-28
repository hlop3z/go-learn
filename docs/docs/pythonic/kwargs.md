# Key-Value Arguments

Go does not have an exact equivalent to Python's `**kwargs`, but you can achieve similar functionality using one of the following approaches:

---

## Dict annotation

```go
type Dict = map[string]interface{}
```

## Kwargs Examples

### 1. **Using Variadic Arguments (`...interface{}`)**

You can use Go's variadic arguments to pass an arbitrary number of arguments. This is the closest equivalent to Python's `*args` or `**kwargs` but does not enforce named arguments.

```go
package main

import "fmt"

// Function that accepts variadic arguments
func printValues(key string, values ...interface{}) {
    fmt.Printf("Key: %s, Values: %v\n", key, values)
}

func main() {
    printValues("example", 1, "text", true, 3.14)
}
```

**Explanation**:

- `values ...interface{}` allows you to pass any number of arguments of any type.
- You must handle the data manually within the function.

---

### 2. **Using Maps for Named Arguments**

To mimic `**kwargs`, you can pass a map of key-value pairs.

```go
package main

import "fmt"

type Dict = map[string]interface{}

// Function that accepts a map for named arguments
func process(kwargs Dict) {
    for key, value := range kwargs {
        fmt.Printf("Key: %s, Value: %v\n", key, value)
    }
}

func main() {
    kwargs := Dict{
        "name":  "Alice",
        "age":   30,
        "admin": true,
    }

    process(kwargs)
}
```

**Explanation**:

- A `map[string]interface{}` acts as a dictionary for named arguments.
- This approach makes it easier to manage key-value pairs.

---

### 3. **Using Structs for Typed Named Arguments**

If the keys and their types are predefined, you can use a struct. This provides compile-time type safety.

```go
package main

import "fmt"

// Struct for named arguments
type Options struct {
    Name  string
    Age   int
    Admin bool
}

// Function that accepts a struct
func greet(opts Options) {
    fmt.Printf("Hello, %s! Age: %d, Admin: %v\n", opts.Name, opts.Age, opts.Admin)
}

func main() {
    opts := Options{Name: "Alice", Age: 30, Admin: true}
    greet(opts)
}
```

**Explanation**:

- A `struct` groups predefined keys with their respective types.
- This is the most idiomatic way in Go when you know the expected arguments.

---

### Comparison with Python `**kwargs`

| Feature              | Python `**kwargs`      | Go Alternatives                 |
| -------------------- | ---------------------- | ------------------------------- |
| Arbitrary arguments  | `**kwargs`             | `...interface{}`                |
| Named arguments      | `**kwargs` (dict-like) | `map[string]interface{}`        |
| Type-safe named args | Not directly possible  | `struct` with predefined fields |

---

### Summary

- Use `...interface{}` for flexibility but without named arguments.
- Use `map[string]interface{}` for dynamic named arguments (closest to `**kwargs`).
- Use a `struct` for type safety when arguments are known ahead of time.

Each approach has its use case depending on the level of flexibility or type safety you need!

## Kwargs and Args

### Using a Slice and a Map

```go
package main

import "fmt"

// Method that mimics *args and **kwargs, returns a reusable function
func method(args []string) func(kwargs map[string]interface{}) map[string]interface{} {
    return func(kwargs map[string]interface{}) map[string]interface{} {
        result := make(map[string]interface{})

        // Iterate over args and populate result with corresponding kwargs
        for _, key := range args {
            if value, exists := kwargs[key]; exists {
                result[key] = value
            }
        }

        return result
    }
}

func main() {
 // Example usage
 args := []string{"name", "age"}
 kwargs := map[string]interface{}{
  "name": "John Doe",
  "age":  30,
  "city": "New York",
 }

 result := method(args, kwargs)
 fmt.Println(result) // Output: map[name:John Doe age:30]
}
```
