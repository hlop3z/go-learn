# GOJA - JavaScript - ECMAScript (5.1)

**File:** `tools_js/add.js`

```js
const add = (data) => a + b;
const emit = (val) => ({ data: val });
```

**File:** `main.go`

```go
package main

import (
 "fmt"
 "log"
 "os"

 "github.com/dop251/goja"
)

func main() {
 filePath := "tools_js/add.js"
 preCode, _ := readFile(filePath)
 runJS(preCode)
}

type Dict = map[string]interface{}

// runJS executes the provided JavaScript code and processes the result
func runJS(preCode string) {
 // Create a new Goja VM instance
 vm := goja.New()

 // Append additional JavaScript code to the provided code
 script := preCode + `emit(add(1, 2))`

 // Execute the script in the Goja VM
 result, is_error := vm.RunString(script)
 if is_error != nil {
  log.Printf("Error executing JavaScript: %v", is_error)
  return
 }

 // Convert the `goja.Value` to a Go `interface{}`
 goValue := result.Export()

 // Convert the result to a map (DictMap) if it's a valid type
 if javascript, ok := goValue.(Dict); ok {
  // Print the resulting map as JSON (formatted as a map)
  fmt.Println("RETURN:", javascript["data"])
 } else {
  log.Println("Result is not a valid DictMap.")
 }
}

// readFile reads the content of a file and returns it as a string
func readFile(filePath string) (string, error) {
 // Read the file contents into a byte slice
 data, is_error := os.ReadFile(filePath)
 if is_error != nil {
  return "", fmt.Errorf("could not read file %s: %v", filePath, is_error)
 }

 // Convert the byte slice to a string and return it
 return string(data), nil
}
```
