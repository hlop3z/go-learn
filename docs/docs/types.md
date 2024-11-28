# Types

Go has a set of built-in types that are divided into basic types, composite types, and more advanced system types. Here's an overview:

---

## **1. Basic Types**

These are the fundamental data types in Go:

| Category       | Type                                          | Description                                       | Example           |
| -------------- | --------------------------------------------- | ------------------------------------------------- | ----------------- |
| Boolean        | `bool`                                        | Represents true/false values.                     | `true`, `false`   |
| Integer        | `int`, `int8`, `int16`, `int32`, `int64`      | Signed integers of various sizes.                 | `42`, `-128`      |
| Unsigned Int   | `uint`, `uint8`, `uint16`, `uint32`, `uint64` | Unsigned integers (non-negative).                 | `0`, `255`        |
| Alias Int      | `byte`                                        | Alias for `uint8`, typically for raw binary data. | `255`             |
| Alias Int      | `rune`                                        | Alias for `int32`, used for Unicode code points.  | `'A'`, `'\u4F60'` |
| Floating Point | `float32`, `float64`                          | Floating-point numbers for decimals.              | `3.14`, `1.618`   |
| Complex        | `complex64`, `complex128`                     | Complex numbers with real and imaginary parts.    | `1 + 2i`          |
| String         | `string`                                      | Immutable sequences of bytes representing text.   | `"Hello, World!"` |

---

## **2. Composite Types**

These are types that combine other types:

| Type          | Description                                         | Example                                         |
| ------------- | --------------------------------------------------- | ----------------------------------------------- |
| **Array**     | Fixed-length sequence of elements of the same type. | `var a [5]int = [5]int{1, 2, 3, 4, 5}`          |
| **Slice**     | Dynamic-length array abstraction.                   | `var s []int = []int{1, 2, 3}`                  |
| **Map**       | Key-value pair collection.                          | `var m map[string]int = map[string]int{"a": 1}` |
| **Struct**    | Collection of named fields of various types.        | `type Point struct { X, Y int }`                |
| **Pointer**   | Reference to a memory address of a variable.        | `var p *int = &x`                               |
| **Function**  | First-class functions with optional return values.  | `func add(a, b int) int { return a + b }`       |
| **Channel**   | Communication mechanism for goroutines.             | `ch := make(chan int)`                          |
| **Interface** | Defines a set of method signatures.                 | `type Reader interface { Read(p []byte) }`      |

---

## **3. System Types**

These types relate to more advanced operations and system-level functionality:

| Type               | Description                                        | Example                         |
| ------------------ | -------------------------------------------------- | ------------------------------- |
| **uintptr**        | An integer large enough to hold a pointer.         | Used for low-level programming. |
| **unsafe.Pointer** | Allows converting between different pointer types. | Requires `unsafe` package.      |

---

## **4. Alias and Special Types**

These types are aliases for specific use cases:

| Type   | Alias For | Description                   | Example |
| ------ | --------- | ----------------------------- | ------- |
| `byte` | `uint8`   | Used for raw data or bytes.   | `'A'`   |
| `rune` | `int32`   | Used for Unicode code points. | `'😊'`  |

---

## **5. Constants and Their Types**

Constants in Go have specific types, including:

| Category          | Example                   |
| ----------------- | ------------------------- |
| Untyped constants | `const Pi = 3.14`         |
| Typed constants   | `const Pi float64 = 3.14` |

---

## **6. Example Code Using Built-in Types**

```go
package main

import "fmt"

func main() {
    // Basic types
    var age int = 25
    var price float64 = 19.99
    var active bool = true
    var name string = "Alice"

    // Composite types
    var numbers = []int{1, 2, 3} // Slice
    var grades = map[string]int{"math": 95, "science": 90} // Map

    // Struct
    type Point struct {
        X int
        Y int
    }
    var p = Point{X: 10, Y: 20}

    // Print values
    fmt.Println(age, price, active, name)
    fmt.Println(numbers, grades)
    fmt.Println(p)
}
```

---

## **Type Summary**

Go's type system is simple yet powerful, combining built-in types for primitive operations with composite types for complex data structures. These types ensure both performance and clarity in code.
