# Go `JSON`

Go provides built-in support for working with the JSON data format through the `encoding/json` package. This includes powerful encoders and decoders for marshaling (converting Go structs to JSON) and unmarshaling (parsing JSON into Go structs).

**File:** `data/input.json`

```json
{
  "name": "John Doe",
  "age": 30,
  "email": "john.doe@example.com",
  "skills": ["Go", "Python", "JavaScript"]
}
```

**File:** `main.go`

```go
package main

import (
 "encoding/json"
 "fmt"
 "log"
 "os"
)

// User struct defines the structure of the JSON data we are working with.
// The struct tags (e.g., `json:"name"`) map JSON field names to struct fields.
type User struct {
 Name   string   `json:"name"`   // User's name
 Age    int      `json:"age"`    // User's age
 Email  string   `json:"email"`  // User's email address
 Skills []string `json:"skills"` // List of user's skills
}

// check is a utility function for error handling. If an error occurs,
// it logs the error message and stops the program.
func check(e error) {
 if e != nil {
  log.Fatalf("Failed with ERROR: %v", e)
  panic(e)
 }
}

// LoadsJson demonstrates how to read and parse JSON data from a file into a Go struct.
func LoadsJson() {
 // Read the JSON file's content.
 // This function reads the file "data/input.json" and returns its contents as a byte slice.
 data, err := os.ReadFile("data/input.json")
 check(err)

 // Declare a variable of type `User` to hold the parsed data.
 var user User

 // Parse the JSON data into the `user` variable.
 // `json.Unmarshal` converts the JSON data into the Go struct based on matching field names.
 err = json.Unmarshal(data, &user)
 check(err)

 // Display the parsed data to the console in a human-readable format.
 fmt.Println("Parsed JSON Data:")
 fmt.Printf("Name: %s\n", user.Name)
 fmt.Printf("Age: %d\n", user.Age)
 fmt.Printf("Email: %s\n", user.Email)
 fmt.Printf("Skills: %v\n", user.Skills)
}

// DumpsJson demonstrates how to create a Go struct, convert it to JSON,
// and write the JSON data to a file.
func DumpsJson() {
 // Create an instance of the `User` struct with sample data.
 user := User{
  Name:   "John Doe",                                // Name of the user
  Age:    30,                                       // Age of the user
  Email:  "john.doe@example.com",                   // Email of the user
  Skills: []string{"Go", "Python", "JavaScript"},   // List of skills
 }

 // Convert the Go struct to a JSON-formatted byte slice with indentation.
 // `json.MarshalIndent` adds formatting to make the output more readable.
 userJSON, err := json.MarshalIndent(user, "", "  ")
 check(err)

 // Write the JSON data to a file called "data/output.json".
 // `0644` sets the file permissions so that the owner can read/write, and others can only read.
 err = os.WriteFile("data/output.json", userJSON, 0644)
 check(err)

 fmt.Println("JSON data successfully written to output.json")
}

// main is the entry point of the program. It calls both `LoadsJson` and `DumpsJson`
// to demonstrate reading from and writing to JSON files.
func main() {
 LoadsJson()
 DumpsJson()
}
```
