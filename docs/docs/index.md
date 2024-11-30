# Go Commands

In **Go**, there are several built-in commands that you'll likely use frequently during development. These commands are part of the **Go Tool** suite, which helps with tasks like building, testing, and managing dependencies. Here’s a list of the most commonly used **Go** commands:

[Learn More . . . GO Documentation](https://go.dev/doc/tutorial/getting-started)

---

## **Summary of Most Common Commands**

- **Get Started**: `go mod init <project/module>` for example: `go mod init example/hello`
- **Development**: `go run`, `go build`, `go install`
- **Dependency Management**: `go get`, `go mod tidy`, `go mod vendor`
- **Testing**: `go test`, `go test -cover`
- **Formatting**: `go fmt`
- **Documentation**: `go doc`
- **Environment**: `go version`, `go env`
- **Clean-up**: `go clean`
- **Advanced**: `go tool pprof`, `go generate`

---

## Commands

These are the core Go commands that you will use frequently to build, test, document, and manage dependencies for your Go applications.

---

### **1. `go run`**

- **Purpose**: Compiles and runs Go source files.
- **Usage**:

  ```bash
  go run main.go
  ```

- **Use case**: When you want to quickly compile and execute a Go program without creating an executable.

---

### **2. `go build`**

- **Purpose**: Compiles Go source code into an executable.
- **Usage**:

  ```bash
  go build
  ```

  To specify a filename:

  ```bash
  go build -o output-file-name main.go
  ```

- **Use case**: Compiling your Go program into a binary for deployment or testing.

---

### **3. `go install`**

- **Purpose**: Compiles and installs the Go binary (places the executable in `$GOPATH/bin` or `$GOBIN`).
- **Usage**:

  ```bash
  go install
  ```

- **Use case**: Installs a compiled program to the Go workspace, making it available globally (useful for tools or utilities).

---

### **4. `go test`**

- **Purpose**: Runs tests defined in the Go codebase.
- **Usage**:

  ```bash
  go test
  ```

  To run tests in a specific file:

  ```bash
  go test -v ./...
  ```

  For benchmarking:

  ```bash
  go test -bench .
  ```

- **Use case**: Running unit tests and benchmarks in your project.

---

### **5. `go fmt`**

- **Purpose**: Automatically formats Go source code according to the Go standard style.
- **Usage**:

  ```bash
  go fmt main.go
  ```

  Or to format all Go files in the current directory and subdirectories:

  ```bash
  go fmt ./...
  ```

- **Use case**: Ensuring your code adheres to Go's formatting conventions.

---

### **6. `go get`**

- **Purpose**: Downloads and installs dependencies from external repositories.
- **Usage**:

  ```bash
  go get github.com/gin-gonic/gin
  ```

- **Use case**: Installing external libraries or modules into your project’s workspace.

---

### **7. `go mod`**

- **Purpose**: Manages dependencies for Go projects using Go Modules.
  - `go mod init` — Initializes a new Go module.
  - `go mod tidy` — Removes unused dependencies.
  - `go mod vendor` — Creates a `vendor` directory with dependencies.
  - `go mod verify` — Verifies that dependencies are consistent with `go.sum`.
- **Usage**:

  ```bash
  go mod init
  go mod tidy
  go mod vendor
  go mod verify
  ```

- **Use case**: Managing project dependencies and ensuring version consistency.

---

### **8. `go doc`**

- **Purpose**: Displays documentation for Go packages or functions.
- **Usage**:

  ```bash
  go doc fmt
  go doc fmt.Println
  ```

- **Use case**: Quickly checking documentation for standard or third-party packages.

---

### **9. `go clean`**

- **Purpose**: Removes build artifacts such as binaries, object files, and cached test results.
- **Usage**:

  ```bash
  go clean
  ```

  To clean specific files:

  ```bash
  go clean -i
  ```

- **Use case**: Cleaning up unnecessary files generated during the build or test process.

---

### **10. `go version`**

- **Purpose**: Displays the current version of Go installed on your system.
- **Usage**:

  ```bash
  go version
  ```

- **Use case**: Checking the installed Go version.

---

### **11. `go env`**

- **Purpose**: Displays Go environment variables.
- **Usage**:

  ```bash
  go env
  ```

- **Use case**: Viewing your Go environment settings, such as `GOPATH` and `GOBIN`.

---

### **12. `go list`**

- **Purpose**: Lists Go packages and their dependencies.
- **Usage**:

  ```bash
  go list ./...
  ```

- **Use case**: Listing all the Go packages in your project or displaying metadata about a specific package.

---

### **13. `go mod vendor`**

- **Purpose**: Copies dependencies into a `vendor` directory, making them available locally.
- **Usage**:

  ```bash
  go mod vendor
  ```

- **Use case**: Managing dependencies within the `vendor` directory for applications that need it (e.g., in isolated environments or with certain deployment strategies).

---

### **14. `go doc -all`**

- **Purpose**: Shows all documentation available for all Go standard packages.
- **Usage**:

  ```bash
  go doc -all
  ```

- **Use case**: Quickly finding all the documentation for a particular module or package.

---

### **15. `go test -cover`**

- **Purpose**: Runs tests and displays test coverage.
- **Usage**:

  ```bash
  go test -cover
  ```

- **Use case**: Checking how much of your code is covered by tests.

---

### **16. `go generate`**

- **Purpose**: Runs code generators (usually used for generating code at compile-time based on annotations).
- **Usage**:

  ```bash
  go generate
  ```

- **Use case**: Executing code generation commands (like generating mock files or auto-generated code) before building the project.

---

### **17. `go tool pprof`**

- **Purpose**: Profiles and analyzes performance of Go programs.
- **Usage**:

  ```bash
  go tool pprof <binary> <profile>
  ```

- **Use case**: Analyzing the CPU or memory usage of a Go program.

---

### **18. `go run` with multiple files**

- **Purpose**: You can also run multiple Go files together.
- **Usage**:

  ```bash
  go run main.go utils.go
  ```

---

### **19. `go install` for Global Tools**

- **Purpose**: Installs Go tools to `$GOPATH/bin` or `$GOBIN`.
- **Usage**:

  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

- **Use case**: Installing Go-based tools globally.

---

### **20. `go list -m`**

- **Purpose**: Lists information about Go modules.
- **Usage**:

  ```bash
  go list -m all
  ```

- **Use case**: Displays module dependencies, useful when working with Go modules.
