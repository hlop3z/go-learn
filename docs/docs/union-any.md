# Union | Any

Go does not have a built-in **union** type as in languages like C or Rust. However, Go provides alternatives for achieving similar functionality through interfaces and type switches. Here are some common patterns:

## Examples

### 1. **Using `interface{}` as a Union Type**

In Go, the empty interface (`interface{}`) can hold any value, effectively serving as a union of all types. However, you'll need to assert or check the type at runtime.

```go
package main

import "fmt"

// Example function that accepts a union of types
func processValue(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %v\n", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    processValue(42)       // Integer
    processValue("hello")  // String
    processValue(true)     // Boolean
}
```

**Explanation**:

- The `interface{}` type allows holding any value.
- Use a **type switch** (`value.(type)`) to handle different types explicitly.

---

### 2. **Using Structs with Tagged Fields**

You can create a struct with fields for different possible types and use one field at a time.

```go
package main

import "fmt"

// Union type using a struct
type Value struct {
    IntVal   *int
    StringVal *string
    BoolVal  *bool
}

func processValue(value Value) {
    if value.IntVal != nil {
        fmt.Printf("Integer: %d\n", *value.IntVal)
    } else if value.StringVal != nil {
        fmt.Printf("String: %s\n", *value.StringVal)
    } else if value.BoolVal != nil {
        fmt.Printf("Boolean: %v\n", *value.BoolVal)
    } else {
        fmt.Println("No value")
    }
}

func main() {
    intValue := 42
    strValue := "hello"
    boolValue := true

    processValue(Value{IntVal: &intValue})
    processValue(Value{StringVal: &strValue})
    processValue(Value{BoolVal: &boolValue})
}
```

**Explanation**:

- A struct acts as a tagged union where only one field is used at a time.
- Use `nil` to indicate which field is active.

---

### 3. **Using Generics (Go 1.18+)**

If you're working with Go 1.18 or later, you can use generics for type-safe unions by defining constraints.

```go
package main

import "fmt"

// Define a generic constraint
type Union interface {
    int | string | bool
}

func processValue[T Union](value T) {
    fmt.Printf("Value: %v (Type: %T)\n", value, value)
}

func main() {
    processValue(42)       // Integer
    processValue("hello")  // String
    processValue(true)     // Boolean
}
```

**Explanation**:

- `Union` defines a constraint for accepted types (`int | string | bool`).
- The function `processValue` is generic and works with any type that satisfies the constraint.

---

### Comparison of Approaches

| Approach            | Advantages                     | Limitations                       |
| ------------------- | ------------------------------ | --------------------------------- |
| `interface{}`       | Flexible, works with any type  | Type-checking only at runtime     |
| Tagged Struct       | Clear type usage               | Verbose, requires manual handling |
| Generics (Go 1.18+) | Type-safe, compile-time checks | Limited to supported Go versions  |

---

### Best Practices

- Use `interface{}` or tagged structs for older Go versions or when flexibility is paramount.
- Use generics for type-safe union types in Go 1.18+ when the union is well-defined.
