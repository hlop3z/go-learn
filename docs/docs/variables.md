# Variables

## Description

In Go, `=` and `:=` are both used for assignment, but they serve different purposes depending on the context. Here's a breakdown of when to use each:

---

### **`:=` (Short Variable Declaration)**

Use `:=` when **declaring and initializing a variable** at the same time. This is typically used within functions.

#### **Key Points**

- Can only be used **inside a function** (not at the package level).
- Automatically determines the variable's type based on the value assigned.
- Declares and assigns the variable in a single step.

#### **Example**

```go
package main

import "fmt"

func main() {
    // Using := to declare and initialize a variable
    x := 10        // x is an int
    message := "Hi" // message is a string

    fmt.Println(x, message)
}
```

---

### **`=` (Assignment Operator)**

Use `=` when **assigning a value to an already declared variable** or when explicitly declaring variables with the `var` keyword.

#### **Key Points**

- Can be used anywhere, including at the package level.
- Does not declare the variable (requires it to already exist).
- Used for reassigning values to existing variables.

#### **Example**

```go
package main

import "fmt"

func main() {
    var x int         // Declare variable x
    x = 20            // Assign a value to x

    var message string
    message = "Hello" // Assign a value to message

    fmt.Println(x, message)
}
```

---

### **When to Use Each**

| Scenario                                    | Use `:=`               | Use `=`                        |
| ------------------------------------------- | ---------------------- | ------------------------------ |
| Declaring and initializing a variable       | ✅                     | ❌                             |
| Reassigning a value to an existing variable | ❌                     | ✅                             |
| Declaring at the package level              | ❌ (not allowed)       | ✅ (with `var` or `const`)     |
| Inside a function                           | ✅ (for new variables) | ✅ (for already declared vars) |

---

### **Example: Combined Use**

```go
package main

import "fmt"

func main() {
    x := 10       // Declare and initialize x
    fmt.Println(x) // Prints 10

    x = 20        // Reassign a new value to x
    fmt.Println(x) // Prints 20

    y := 30       // Declare and initialize y
    fmt.Println(y) // Prints 30

    y = x + 10    // Reassign y using x
    fmt.Println(y) // Prints 30
}
```

---

### **Pitfalls with `:=`**

1. **Shadowing**: Using `:=` can accidentally declare a new variable instead of updating an existing one.

   ```go
   package main

   import "fmt"

   func main() {
       x := 10
       if true {
           x := 20 // This creates a new variable shadowing the outer x
           fmt.Println(x) // Prints 20 (inner x)
       }
       fmt.Println(x) // Prints 10 (outer x)
   }
   ```

   To avoid this, use `=` when modifying an existing variable.

2. **Cannot Use at the Package Level**:

   ```go
   // This will not compile:
   x := 10 // Error: short variable declaration not allowed at package level
   ```

---

### **Best Practices**

- Use `:=` for new variables when you're writing concise code inside functions.
- Use `=` when updating variables or declaring them explicitly with `var`.
- Avoid shadowing by carefully managing `:=` within nested scopes.
