# Pythonic String Search

```go
package main

import (
 "fmt"
 "strings"
)

func main() {
 text := "Hello, World!"
 prefix := "Hello"
 suffix := "World!"
 substring := "lo, Wo"

 // StartsWith
 if strings.HasPrefix(text, prefix) {
  fmt.Printf("The text starts with '%s'\n", prefix)
 }

 // EndsWith
 if strings.HasSuffix(text, suffix) {
  fmt.Printf("The text ends with '%s'\n", suffix)
 }

 // Contains
 if strings.Contains(text, substring) {
  fmt.Printf("The text contains '%s'\n", substring)
 }

 // Regex
 startsWithPattern := `^Hello`
 startsWithRegex := regexp.MustCompile(startsWithPattern)
 if startsWithRegex.MatchString(text) {
  fmt.Printf("The text starts with 'Hello'\n")
 }
}
```
