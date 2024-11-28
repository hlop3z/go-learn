# Strings

The `fmt.Sprintf` function in Go is part of the `fmt` package and is used to format strings. Unlike `fmt.Printf`, which prints the formatted string directly to standard output, `fmt.Sprintf` returns the formatted string, allowing you to assign it to variables or use it in further operations.

## **Basic Syntax**

```go
func Sprintf(format string, a ...interface{}) string
```

- **`format`**: A string containing format specifiers (placeholders) that define how the provided values should be formatted.
- **`a`**: A variadic parameter for the values to be formatted and substituted into the `format` string.

---

### **Common Format Specifiers**

1. **General Specifiers**:

   - `%v`: Default format for any value.
   - `%+v`: Includes field names for structs.
   - `%#v`: Go syntax representation of the value.
   - `%T`: Type of the value.

2. **Integer Formatting**:

   - `%d`: Decimal (base 10).
   - `%b`: Binary (base 2).
   - `%o`: Octal (base 8).
   - `%x`: Hexadecimal (base 16, lowercase).
   - `%X`: Hexadecimal (base 16, uppercase).

3. **Floating-Point and Complex Numbers**:

   - `%f`: Decimal point (e.g., 123.456).
   - `%e`: Scientific notation (lowercase, e.g., 1.234560e+02).
   - `%E`: Scientific notation (uppercase, e.g., 1.234560E+02).
   - `%.nf`: Specify precision to `n` decimal places.

4. **Strings and Characters**:

   - `%s`: String.
   - `%q`: Quoted string with escape characters.
   - `%c`: Unicode character.
   - `%U`: Unicode code point (e.g., `U+1234`).

5. **Booleans**:

   - `%t`: Prints `true` or `false`.

6. **Pointer Addresses**:
   - `%p`: Pointer address.

---

### **Examples**

#### **Basic Formatting**

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30

    result := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(result) // Output: Name: Alice, Age: 30
}
```

---

#### **Using Structs**

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Bob", 25}

    result := fmt.Sprintf("Struct: %v", p)
    fmt.Println(result) // Output: Struct: {Bob 25}

    resultWithFields := fmt.Sprintf("Struct with fields: %+v", p)
    fmt.Println(resultWithFields) // Output: Struct with fields: {Name:Bob Age:25}

    goSyntax := fmt.Sprintf("Go syntax: %#v", p)
    fmt.Println(goSyntax) // Output: Go syntax: main.Person{Name:"Bob", Age:25}
}
```

---

#### **Formatting Numbers**

```go
package main

import "fmt"

func main() {
    num := 255

    binary := fmt.Sprintf("Binary: %b", num)
    fmt.Println(binary) // Output: Binary: 11111111

    octal := fmt.Sprintf("Octal: %o", num)
    fmt.Println(octal) // Output: Octal: 377

    hex := fmt.Sprintf("Hexadecimal: %x", num)
    fmt.Println(hex) // Output: Hexadecimal: ff

    upperHex := fmt.Sprintf("Uppercase Hex: %X", num)
    fmt.Println(upperHex) // Output: Uppercase Hex: FF
}
```

---

#### **Floating-Point Precision**

```go
package main

import "fmt"

func main() {
    pi := 3.14159265359

    result := fmt.Sprintf("Pi (default): %f", pi)
    fmt.Println(result) // Output: Pi (default): 3.141593

    resultPrecision := fmt.Sprintf("Pi (2 decimal places): %.2f", pi)
    fmt.Println(resultPrecision) // Output: Pi (2 decimal places): 3.14
}
```

---

#### **Combining Multiple Types**

```go
package main

import "fmt"

func main() {
    name := "Eve"
    isStudent := true
    grade := 89.5

    result := fmt.Sprintf("Name: %s, Student: %t, Grade: %.1f", name, isStudent, grade)
    fmt.Println(result) // Output: Name: Eve, Student: true, Grade: 89.5
}
```

---

### **Tips and Best Practices**

1. **Readability**: Use meaningful format specifiers to enhance the readability of your code.
2. **Escaping `%`**: To include a literal `%` in the output, use `%%`.

   ```go
   fmt.Sprintf("Percentage: 50%%") // Output: Percentage: 50%
   ```

3. **Error Logging**: Combine `Sprintf` with error handling to create informative error messages.

   ```go
   errMsg := fmt.Sprintf("Error: Unable to find user with ID %d", userID)
   log.Println(errMsg)
   ```

4. **Debugging**: Use `%#v` to inspect complex objects during debugging.

---

With `fmt.Sprintf`, you have a versatile tool for formatting strings dynamically in Go. It is widely used for logging, error handling, and constructing strings for user-facing outputs.
