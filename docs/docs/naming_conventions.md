# Naming Conventions

In Go, naming conventions are consistent and idiomatic, helping to make code readable and maintainable. Here's an overview of the recommended naming conventions in Go:

## Summary Table

| Element        | Convention   | Example                  |
| -------------- | ------------ | ------------------------ |
| Files          | `snake_case` | `user_service.go`        |
| Packages       | `lowercase`  | `http`, `json`           |
| Functions      | `camelCase`  | `getUserName`, `update`  |
| Variables      | `camelCase`  | `userName`, `maxRetries` |
| Exported Funcs | `PascalCase` | `GetUserName`            |
| Constants      | `PascalCase` | `MaxConnections`         |
| Structs        | `PascalCase` | `User`, `Product`        |
| Interfaces     | `PascalCase` | `Reader`, `Logger`       |

---

## Additional Notes

1. **Exporting Symbols**: If something starts with an uppercase letter, it becomes exported and accessible outside the package.
2. **Keep It Short**: Go values simplicity, so avoid verbose names.
3. **Acronyms**: Use consistent casing for acronyms like `HTTP`, `URL`, `ID`.
   - Example: `HTTPRequest`, `userID`.

By following these conventions, your Go code will remain idiomatic and align well with community standards.

---

## Examples

### **1. Functions and Methods**

- Use **camelCase** for function and method names.
- Exported functions/methods (those that are accessible outside the package) must start with an **uppercase letter** (PascalCase).
- Non-exported functions/methods (internal to the package) start with a **lowercase letter** (camelCase).

#### Example

```go
// Exported function (visible outside the package)
func GetUserName() string {
    return "John Doe"
}

// Non-exported function (only accessible within the package)
func calculateSum(a, b int) int {
    return a + b
}
```

---

### **2. Variables and Constants**

- Use **camelCase** for variables.
- Use **ALL_CAPS_WITH_UNDERSCORES** or **PascalCase** for constants, though PascalCase is more common in Go.

#### Example

```go
// Variables
var userName string
var maxRetries int

// Constants
const MaxConnections = 10
const DefaultTimeout = 30
```

---

### **3. Structs**

- Use **PascalCase** for struct names (always exported by convention).

#### Example

```go
type User struct {
    ID   int
    Name string
}
```

---

### **4. Interfaces**

- Use **PascalCase** for interface names.
- Interface names should describe behavior and often end in **-er** (e.g., `Reader`, `Writer`, `Logger`) or have a descriptive name like `Database`.

#### Example

```go
type Reader interface {
    Read(data []byte) (int, error)
}
```

---

### **5. Packages**

- Use **lowercase letters** for package names.
- Avoid underscores, hyphens, or mixed case.
- Package names should be short and describe their functionality.

#### Example

```plaintext
math, strings, http, json
```

When importing:

```go
import "github.com/username/awesome_module/package_one"
```

---

### **6. Files**

- Use **snake_case** for file names.
- File names should reflect the contents or primary functionality.

#### Example

```plaintext
user_service.go
database_connection.go
http_handler.go
```

---

### **7. Test Functions**

- Test functions should start with `Test` followed by a descriptive name in **PascalCase**.
- Benchmarks and examples follow similar conventions.

#### Example

```go
func TestCalculateSum(t *testing.T) {
    result := calculateSum(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```
