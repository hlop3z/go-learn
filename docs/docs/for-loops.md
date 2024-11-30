# For Loop Blocks

In Go, the **`for` loop** is the only looping construct available. However, it is versatile enough to cover the functionality of `while` and `foreach` loops found in other languages. Below are examples to illustrate different ways you can use a `for` loop in Go:

---

## Examples

### 1. **Basic `for` Loop** (like Python's `for i in range(n)`)

```go
package main

import "fmt"

func main() {
    // Loop from 0 to 4
    for i := 0; i < 5; i++ {
        fmt.Println("Value of i:", i)
    }
}
```

**Explanation**:

- `i := 0` initializes the loop variable.
- `i < 5` is the condition to continue the loop.
- `i++` increments `i` after each iteration.

---

### 2. **Infinite Loop** (like Python's `while True`)

```go
package main

import "fmt"

func main() {
    count := 0
    for {
        fmt.Println("Infinite loop iteration:", count)
        count++
        if count >= 5 { // Break the loop after 5 iterations
            break
        }
    }
}
```

**Explanation**:

- The loop runs indefinitely unless you explicitly use `break` to exit.

---

### 3. **Conditional `for` Loop** (like Python's `while` loop)

```go
package main

import "fmt"

func main() {
    count := 0
    for count < 5 { // Continue as long as count is less than 5
        fmt.Println("Count is:", count)
        count++
    }
}
```

**Explanation**:

- This behaves like a `while` loop in other languages.

---

### 4. **Iterating Over Slices/Arrays**

```go
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50}

    // Using for loop with index
    for i, num := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
}
```

**Explanation**:

- `range` returns both the index and the value for each element in the slice.
- If you only need the value, use `_` to ignore the index.

---

### 5. **Iterating Over Maps**

```go
package main

import "fmt"

func main() {
    ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}

    for name, age := range ages {
        fmt.Printf("%s is %d years old.\n", name, age)
    }
}
```

**Explanation**:

- `range` works for maps and gives you the key-value pair.

---

### 6. **Using `continue` and `break`**

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        if i > 7 {
            break // Exit the loop if i > 7
        }
        fmt.Println("Odd number:", i)
    }
}
```

**Explanation**:

- `continue` skips to the next iteration.
- `break` exits the loop entirely.

---

### Key Features

- **Initialization and Increment**: Can be omitted (e.g., `for ; i < 5; ;`).
- **Infinite Loops**: Use `for` without conditions.
- **Iterating Collections**: Use `range`.

With these variations, Go's `for` loop is both simple and powerful.
