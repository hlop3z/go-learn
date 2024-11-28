# DateTime

```go
package main

import (
    "fmt"
    "time"
)

// Always use UTC
func DateTimeUTC() time.Time {
    return time.Now().UTC()
}

// Test DateTime
func main() {
    fmt.Println("Current UTC Time:", DateTimeUTC())
}
```
