# Command-Line Interface

If you're looking for a Go solution similar to Python's `argparse` for handling command-line arguments, Go has several libraries that help with parsing and handling command-line arguments. The standard `flag` package provides basic functionality, but if you're looking for something closer to `argparse` with more features like subcommands, argument types, and
default values, libraries like `cobra` and `kingpin` are good options.

## Examples

### 1. **Using the Go Standard Library: `flag`**

Go’s built-in `flag` package is simple and provides basic functionality for parsing command-line flags. It’s a good starting point for simple scripts.

#### Example Using `flag`

```go
package main

import (
 "flag"
 "fmt"
)

func main() {
 // Define flags
 name := flag.String("name", "World", "Name to greet")
 age := flag.Int("age", 0, "Your age")
 help := flag.Bool("help", false, "Show Help")

 // Parse the command-line flags
 flag.Parse()

 // Use the parsed flags
 if *help {
  fmt.Printf("Help: %v", *help)
 } else {
  fmt.Printf("Hello, %s!\n", *name)
  fmt.Printf("You are %d years old.\n", *age)
 }
}
```

Usage:

```bash
go run main.go --name John --age 30
```

This would output:

```bash
Hello, John!
You are 30 years old.
```

### 2. **Using `cobra` for More Advanced Features**

[`cobra`](https://github.com/spf13/cobra) is a more feature-rich library for building command-line applications in Go. It supports commands, subcommands, argument parsing, flags, and more. It’s a great choice for building applications similar to Python's `argparse`.

#### Example Using `cobra`

```go
package main

import (
 "fmt"
 "github.com/spf13/cobra"
)

func main() {
 var name string
 var age int

 // Define the root command
 var rootCmd = &cobra.Command{
  Use:   "greet",
  Short: "A simple greeting app",
  Run: func(cmd *cobra.Command, args []string) {
   fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
  },
 }

 // Define flags for the root command
 rootCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
 rootCmd.Flags().IntVarP(&age, "age", "a", 0, "Age of the person")

 // Execute the root command
 if err := rootCmd.Execute(); err != nil {
  fmt.Println(err)
 }
}
```

Usage:

```bash
go run main.go greet -n John -a 30
```

This would output:

```bash
Hello, John! You are 30 years old.
```

### 3. **Go + `os/exec` for Running Other Scripts**

Go can execute external scripts and commands using the `os/exec` package. This allows you to run shell scripts, Python scripts, or other programs from within your Go program.

#### Example: Running a Shell Script from Go

```go
package main

import (
 "fmt"
 "os/exec"
)

func main() {
 // Run a shell command (in this case, a simple "ls" command)
 cmd := exec.Command("ls", "-l")
 output, err := cmd.CombinedOutput()
 if err != nil {
  fmt.Println("Error executing command:", err)
  return
 }

 // Print the output of the command
 fmt.Println(string(output))
}
```

### Conclusion

For Go scripting similar to Python's `argparse`, you can:

1. Use Go’s built-in `flag` package for basic flag parsing.
2. Use `cobra` for advanced features, including subcommands and dynamic argument parsing.
3. Use `os/exec`: To run external scripts or commands.
