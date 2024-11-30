# Read and Write from/to Files

In Go, you can read from and write to text files using the `os` and `bufio` packages (for more control). Below are examples of how to read from and write to text files in Go.

## Key Points

- `os.ReadFile`: Reads the entire file into memory (ideal for smaller files).
- `bufio.Scanner`: Reads the file line-by-line, which is memory efficient for larger files.
- `os.WriteFile`: Writes the content to a file, overwriting any existing content.
- `os.OpenFile`: Opens a file and provides more control, such as appending content.

These examples should give you a solid foundation for working with text files in Go.

### Reading a Text File

Here’s how you can read a text file in Go:

#### 1. Using `os.ReadFile` (Go 1.16+)

This is the simplest and recommended way to read an entire file in Go.

```go
package main

import (
 "fmt"
 "os"
)

func main() {
 // Reading file content
 data, err := os.ReadFile("example.txt")
 if err != nil {
  fmt.Println("Error reading file:", err)
  return
 }

 // Print the content of the file
 fmt.Println(string(data))
}
```

#### 2. Using `bufio.Scanner` (More Control)

If you need to read the file line by line, you can use the `bufio` package to scan through the file.

```go
package main

import (
 "bufio"
 "fmt"
 "os"
)

func main() {
 file, err := os.Open("example.txt")
 if err != nil {
  fmt.Println("Error opening file:", err)
  return
 }
 defer file.Close()

 // Create a scanner to read the file line by line
 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
  fmt.Println(scanner.Text()) // Print each line
 }

 if err := scanner.Err(); err != nil {
  fmt.Println("Error reading file:", err)
 }
}
```

### Writing to a Text File

To write to a text file, you can use the `os` package for opening/creating files and `io` or `bufio` for writing.

#### 1. Using `os.WriteFile` (Go 1.16+)

This method is simple and suitable when you want to overwrite the content of the file.

```go
package main

import (
 "fmt"
 "os"
)

func main() {
 content := []byte("Hello, this is a sample text file.")

 // Writing to the file
 err := os.WriteFile("output.txt", content, 0644)
 if err != nil {
  fmt.Println("Error writing file:", err)
  return
 }

 fmt.Println("File written successfully!")
}
```

#### 2. Using `bufio.Writer` (More Control)

If you need more control, such as appending data or writing line by line, you can use the `bufio.Writer`:

```go
package main

import (
 "bufio"
 "fmt"
 "os"
)

func main() {
 // Open the file in write mode, create it if it doesn't exist
 file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
 if err != nil {
  fmt.Println("Error opening file:", err)
  return
 }
 defer file.Close()

 // Create a buffered writer
 writer := bufio.NewWriter(file)

 // Write data to the file
 _, err = writer.WriteString("Hello, this is another line of text!\n")
 if err != nil {
  fmt.Println("Error writing to file:", err)
  return
 }

 // Make sure to flush the buffer to the file
 writer.Flush()

 fmt.Println("Text written to file successfully!")
}
```
