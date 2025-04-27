# Using the (`*Pointer` or `Value`)

In Go, the `&` operator is used to get the **address** of a variable. When it comes to types, using `&` with a type generally appears in the context of initializing or passing pointers to types. Here are the common scenarios where `&` is used:

---

## 1. **Using `*AccessControl` (Pointer Receiver)**

### Code 1

```go
func (ac *AccessControl) SetRules(rules Rules) *AccessControl {
 ac.Rules = rules
 return ac
}
```

### When to Use 1

1. **Modifying the Receiver's State**:

   - Use a pointer receiver when the method modifies the state of the struct it operates on.
   - In this case, `ac.Rules = rules` directly updates the `Rules` field of the original `AccessControl` instance.

2. **Avoiding Copying Large Structures**:

   - If the `AccessControl` struct is large, passing it by pointer avoids the overhead of copying it for each method call.

3. **Chainable Methods**:

   - Returning a pointer (`*AccessControl`) allows for method chaining:

     ```go
     ac.SetRules(rules).AnotherMethod()
     ```

---

## 2. **Using `AccessControl` (Value Receiver)**

### Code 2

```go
func (ac AccessControl) SetRules(rules Rules) AccessControl {
 ac.Rules = rules
 return ac
}
```

### When to Use 2

1. **Immutable Operations**:

   - Use a value receiver if the method doesn't need to modify the receiver's state.
   - The `ac.Rules = rules` change affects only a copy of the struct and doesn't alter the original instance.

2. **Small Structs**:

   - If the struct is small, passing by value is efficient and doesn't introduce pointer dereferencing overhead.

3. **Functional Programming Style**:

   - This approach supports a functional programming style where the original object remains unchanged, and a modified copy is returned:

     ```go
     newAC := ac.SetRules(rules)
     ```

---

### Key Differences

| **Aspect**             | **Pointer Receiver (`*AccessControl`)** | **Value Receiver (`AccessControl`)** |
| ---------------------- | --------------------------------------- | ------------------------------------ |
| **State Modification** | Modifies the original instance          | Modifies a copy of the instance      |
| **Performance**        | Avoids copying large structs            | Copies the struct                    |
| **Method Chaining**    | Supports chaining                       | Chaining requires extra assignment   |
| **Immutability**       | Mutates original                        | Maintains immutability               |

---

### Practical Recommendation

- **If the method is designed to modify the struct (e.g., setters like `SetRules`), use a pointer receiver (`*AccessControl`).**
- **If the method doesn't modify the struct and is meant to work on a copy or return a new version, use a value receiver (`AccessControl`).**

In your case, since `SetRules` is clearly a setter and modifies the `Rules` field, the pointer receiver version is the more appropriate choice.

---

## Examples

### 1. **Creating Pointers to Structs**

When you use `&` before a type, you are creating a pointer to that type. For example:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a pointer to a struct
    p := &Person{Name: "Alice", Age: 30}
    fmt.Println(p)        // Output: &{Alice 30}
    fmt.Println(p.Name)   // Accessing fields via the pointer
}
```

- **When to use**:
  - When you want to work with a pointer directly instead of a value.
  - To avoid copying the entire struct when passing it around.

---

### 2. **Passing Pointers to Functions**

When a function expects a pointer to a type, you use `&` to pass the address of a value.

```go
package main

import "fmt"

func updateAge(p *Person) {
    p.Age = 35
}

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Bob", Age: 30}
    updateAge(&person) // Passing a pointer to the struct
    fmt.Println(person.Age) // Output: 35
}
```

- **When to use**:
  - When you need the function to modify the original value (pass-by-reference behavior).
  - When working with large data structures to avoid copying them.

---

### 3. **Using `new` for Pointer Initialization**

You can use `new` to create pointers, but you can also use `&` with a type literal for clarity:

```go
package main

import "fmt"

type Counter struct {
    Value int
}

func main() {
    // Using & to create a pointer
    c := &Counter{Value: 1}
    c.Value++
    fmt.Println(c.Value) // Output: 2
}
```

- **When to use**:
  - Use `&` for explicit pointer creation alongside type initialization.
  - Use `new` when no fields need initialization (`new(T)` returns `*T`).

---

### 4. **Dereferencing and Modifying Values**

You use `&` to create a pointer and `*` to dereference it:

```go
package main

import "fmt"

func main() {
    x := 10
    ptr := &x // Pointer to x
    *ptr = 20 // Dereferencing the pointer to modify x
    fmt.Println(x) // Output: 20
}
```

- **When to use**:
  - Use `&` when you want to reference a variable.
  - Use `*` to modify or access the value pointed to by the pointer.

---

### Summary

- **Use `&`** when:

  - You need a pointer to a type (e.g., `&Struct{}`).
  - Passing variables by reference to avoid copying or to allow modification.
  - Initializing a pointer to a type inline.

- **Don't use `&`** if:
  - You are fine with working with a copy of the value.
  - The function does not require a pointer.

---

## Comparison

When deciding between using pointers (e.g., `&Struct{}`) or values (e.g., `Struct{}`), the most **efficient approach** depends on several factors, including the size of the data, frequency of copying, and the specific use case. Here's a breakdown of efficiency considerations:

---

### 1. **Using Pointers (`&Struct{}`)**

#### Advantages

- **Avoids Copying Large Data**:

  - When working with large structs, passing a pointer avoids copying the entire struct, which saves memory and CPU overhead.
  - For example:

    ```go
    type LargeStruct struct {
        Field1 [1024]byte
        Field2 int
    }

    // Passing by value copies the entire structure
    func processByValue(data LargeStruct) {}

    // Passing by pointer avoids copying
    func processByPointer(data *LargeStruct) {}
    ```

- **Enables Mutability**:

  - Functions receiving pointers can modify the original data, avoiding the need to return modified copies.
  - Useful in scenarios like in-place updates, caching, or synchronization primitives.

- **Consistency for Shared Data**:
  - If multiple parts of the program work on the same data, pointers ensure that changes are reflected everywhere.

#### Disadvantages

- **Potential for Nil Pointers**:

  - Pointers can be `nil`, leading to runtime panics if not checked properly.
  - Example:

    ```go
    var p *Person
    fmt.Println(p.Name) // Causes panic
    ```

- **Dereferencing Overhead**:

  - Accessing data through a pointer incurs a slight overhead due to dereferencing.

- **Garbage Collector Pressure**:
  - Using pointers can increase the load on Go's garbage collector because the collector tracks heap allocations and pointers.

---

### 2. **Using Values (`Struct{}`)**

#### Advantages

- **Simpler Code**:

  - Avoids nil pointer dereference issues and makes code easier to reason about.
  - Example:

    ```go
    person := Person{Name: "Alice", Age: 30}
    fmt.Println(person.Name) // Direct access, no pointer needed
    ```

- **Efficient for Small Data**:

  - Passing small structs (e.g., 2–3 fields) by value is efficient because copying them is inexpensive compared to the cost of managing pointers.
  - Example:

    ```go
    type SmallStruct struct {
        X, Y int
    }
    ```

- **Stack Allocation**:
  - Value types are typically allocated on the stack, reducing garbage collector pressure.

#### Disadvantages

- **Expensive for Large Data**:

  - Copying large structs repeatedly can lead to performance degradation.
  - Example:

    ```go
    type LargeStruct struct {
        Field1 [1024]byte
        Field2 int
    }
    ```

- **Immutability**:
  - Functions working with values cannot modify the original data, leading to potentially unnecessary copies if changes are needed.

---

### 3. **Which Is More Efficient?**

#### General Guidelines

1. **Use Pointers (`&Struct{}`)** when:

   - The struct is large, or copying it would be expensive.
   - You need to modify the original data in functions.
   - The data is shared between multiple parts of the program.
   - You are creating or working with linked data structures (e.g., trees, graphs).

2. **Use Values (`Struct{}`)** when:
   - The struct is small, and copying is inexpensive.
   - Immutability is acceptable or desirable.
   - You want to avoid the complexity of managing pointers and nil checks.
   - The struct will be short-lived and doesn't need to persist across function calls.

---

### 4. **Real-World Benchmarking**

Benchmarking is the best way to decide. Here's an example of how you can test both approaches in your specific use case:

```go
package main

import (
 "testing"
)

type LargeStruct struct {
 Field1 [1024]byte
 Field2 int
}

func processByValue(data LargeStruct) {
 data.Field2 = 42
}

func processByPointer(data *LargeStruct) {
 data.Field2 = 42
}

func BenchmarkByValue(b *testing.B) {
 data := LargeStruct{}
 for i := 0; i < b.N; i++ {
  processByValue(data)
 }
}

func BenchmarkByPointer(b *testing.B) {
 data := &LargeStruct{}
 for i := 0; i < b.N; i++ {
  processByPointer(data)
 }
}
```

Run this benchmark using `go test -bench=.`, and you'll see the performance difference between the two approaches.

---

### Summary Comparison

- For **small structs** and **immutable scenarios**, use **values** for simplicity and performance.
- For **large structs**, **mutable data**, or **shared state**, use **pointers** to minimize copying and enable in-place updates.
- When in doubt, benchmark!
