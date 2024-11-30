# IF Statement Blocks

In Go, the `if` statement is similar to other programming languages. It allows you to execute a block of code based on a condition. Here's a detailed guide with examples:

---

## Examples

### **Basic Syntax**

```go
if condition {
    // Code to execute if condition is true
}
```

---

### **Example 1: Basic `if` Statement**

```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}
```

**Output**:

```plaintext
x is greater than 5
```

---

### **Example 2: `if-else` Statement**

```go
package main

import "fmt"

func main() {
    x := 3

    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is 5 or less")
    }
}
```

**Output**:

```plaintext
x is 5 or less
```

---

### **Example 3: `if-else if-else` Statement**

```go
package main

import "fmt"

func main() {
    x := 7

    if x > 10 {
        fmt.Println("x is greater than 10")
    } else if x > 5 {
        fmt.Println("x is greater than 5 but less than or equal to 10")
    } else {
        fmt.Println("x is 5 or less")
    }
}
```

**Output**:

```plaintext
x is greater than 5 but less than or equal to 10
```

---

### **Example 4: Short Variable Declaration in `if`**

You can declare and initialize a variable directly within the `if` statement, and its scope is limited to that block.

```go
package main

import "fmt"

func main() {
    if y := 15; y > 10 { // Declare and initialize y in the if statement
        fmt.Println("y is greater than 10")
    }
    // fmt.Println(y) // Error: y is not defined here
}
```

**Output**:

```plaintext
y is greater than 10
```

---

### **Example 5: Checking Multiple Conditions**

Use logical operators (`&&`, `||`) to combine conditions.

```go
package main

import "fmt"

func main() {
    x := 8

    if x > 5 && x < 10 {
        fmt.Println("x is between 5 and 10")
    }

    if x < 5 || x > 10 {
        fmt.Println("x is not between 5 and 10")
    }
}
```

**Output**:

```plaintext
x is between 5 and 10
```

---

### **Example 6: Comparing Strings**

You can compare strings lexicographically using `if`.

```go
package main

import "fmt"

func main() {
    name := "Alice"

    if name == "Alice" {
        fmt.Println("Hello, Alice!")
    } else {
        fmt.Println("You're not Alice!")
    }
}
```

**Output**:

```plaintext
Hello, Alice!
```

---

### Key Notes

1. **Braces are mandatory**: Unlike some languages (e.g., Python), Go requires braces `{}` for `if` blocks, even for a single statement.
2. **No parentheses**: Parentheses around the condition are optional but discouraged (`if (x > 5)` is valid, but idiomatic Go uses `if x > 5`).
3. **Short variable declaration**: Variables declared in `if` (e.g., `y := 15`) are scoped to the `if` block and its associated `else` or `else if` blocks.

---

### Common Mistakes

- Forgetting braces `{}`.
- Trying to use variables declared in the `if` statement outside its scope.
- Misusing logical operators (`&&`, `||`).

With these examples and tips, you should be able to use `if` effectively in Go!
