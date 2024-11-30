# Go `HTML` Template Rendering

Go provides built-in support for working with HTML content through the `html/template` package. This package helps to separate logic from presentation, allowing dynamic content to be safely inserted into HTML pages. It is especially useful for generating web pages based on templates and user data.

## Template Example

### **Template File:** `templates/index.html`

The following is an example of an HTML template that contains placeholders for dynamic content. These placeholders will be replaced at runtime using data passed from your Go code.

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1.0, user-scalable=no, maximum-scale=1, minimum-scale=1"
    />
    <title>{{.Meta.Title}}</title>
  </head>

  <body>
    <h1>{{.Heading}}</h1>
    <p>{{.Content}}</p>

    <!-- Array/Slice -->
    <ul>
      {{range .Items}}
      <li>{{.}}</li>
      {{end}}
    </ul>

    <!-- IF-Statement -->
    {{if .User.IsAuthenticated}}
    <h1>Welcome back, {{.User.Username}}!</h1>
    {{else}}
    <h1>Welcome, Guest!</h1>
    {{end}}

    <!-- Switch -->
    {{if eq .User.Role.Name "admin"}}
    <h1>Admin Panel</h1>
    {{else if eq .User.Role.Name "user"}}
    <h1>User Dashboard</h1>
    {{else}}
    <h1>Access Denied</h1>
    {{end}}
  </body>
</html>
```

In this template:

- `{{.Heading}}`, `{{.Content}}` and `{{.ETC...}}` are placeholders that will be replaced by data passed to the template when rendering it.

### **Go File:** `main.go`

This Go code reads the template file, creates data to populate the placeholders, and renders the final HTML.

```go
package main

import (
 "bytes"
 "fmt"
 "html/template"
 "os"
)

// Dict is a type alias for a map with string keys and values of any type.
// It is used to pass a flexible data structure to templates.
type Dict = map[string]interface{}

// List is a type alias for a slice of interface{}, typically used for handling ordered collections of mixed types in templates.
type List = []interface{}

func main() {
 // Define a template string by reading the contents of an HTML template file.
 var tmpl = readFile("templates/index.html")

 // Create a User instance with default values.
 user := User{
  IsAuthenticated: true,
  Username:        "John",
  Role:            Role{Name: "user"},
 }

 // Create a data structure to hold dynamic values for rendering.
 // These values will be passed to the template during execution.
 data := Dict{
  "User": user, // Pass the user object to the template.
  "Heading": "Welcome to Go Templates!",                                                                               // Page heading.
  "Content": template.HTML("<strong>This is a dynamic HTML page generated using the html/template package.</strong>"), // Dynamic content to render.
  "Items":   List{"Item 1", "Item 2", "Item 3"},                                                                       // A list of items to display.
  "Meta": Dict{
   "Title": "My Go HTML Template", // Dynamic metadata for the page.
  },
 }

 // Use the HTML helper function to render the template.
 render := HTML(tmpl) // Get a function to render the template.
 html := render(data) // Call the render function with the data.
 fmt.Println(html)    // Output the rendered HTML to the console.
}

// HTML takes a template string and returns a function that renders the template with the given data.
// This allows for easy reuse of the template rendering logic with different sets of data.
func HTML(tmpl string) func(data interface{}) string {
 // Parse the template string into a Template object.
 t, err := template.New("webpage").Parse(tmpl)
 if err != nil {
  // Panic if there's an error parsing the template.
  panic(err)
 }

 // Return a function that renders the template with the provided data.
 return func(data interface{}) string {
  // Create a buffer to capture the output of the template rendering.
  var buf bytes.Buffer
  // Execute the template, passing in the data.
  err = t.Execute(&buf, data)
  if err != nil {
   // Panic if there's an error executing the template.
   panic(err)
  }

  // Return the generated HTML as a string.
  return buf.String()
 }
}

// readFile reads the contents of a file at the given path and returns it as a string.
// It panics if there is an error reading the file.
func readFile(filePath string) string {
 data, err := os.ReadFile(filePath)
 if err != nil {
  // Panic if there's an error reading the file.
  panic(err)
 }
 return string(data)
}

// Role defines the structure for a user's role in the application.
// It contains an ID, Name, and Level to indicate the role's priority.
type Role struct {
 ID    uint   // Unique identifier for the role.
 Name  string // Name of the role.
 Level uint   // The level or rank of the role.
}

// User represents a user in the application.
// It includes personal details, role, and flags for authentication and account status.
type User struct {
 ID uint // Unique identifier for the user.

 // Role and account information
 Role     Role   // Role associated with the user.
 Username string // Username for the user.
 Email    string // Email address of the user.
 Password string // User's password (hashed for security).

 // Personal details
 FirstName  string // User's first name.
 MiddleName string // User's middle name (optional).
 LastName   string // User's last name.

 // Flags indicating the user's status
 IsStaff         bool // Whether the user is a staff member.
 IsActive        bool // Whether the user is active.
 IsSuperuser     bool // Whether the user is a superuser.
 IsAuthenticated bool // Whether the user is authenticated.

 // DateTime fields to track events for the user
 CreatedAt string // Timestamp for when the user was created (using string for simplicity, ideally time.Time).
 UpdatedAt string // Timestamp for the last update to the user's information.
 DeletedAt string // Timestamp for when the user was deleted (if applicable).
}

```

### Key Concepts

- **Templates**: Templates are HTML files with placeholders (e.g., `{{.Title}}`) that can be dynamically replaced with data.
- **Data**: The data is passed to the template as a map (of type `Dict`) that maps variable names (e.g., `Title`, `Heading`, `Content`) to their values.
- **Rendering**: The template is parsed and executed with the provided data to generate the final HTML content.
- **HTML Safety**: The `template.HTML` type ensures that HTML content is not automatically escaped when rendering.

### How It Works

1. **Template Parsing**: The `html/template` package is used to parse the `index.html` file.
2. **Dynamic Data Insertion**: The placeholders in the HTML template (e.g., `{{.Title}}`) are replaced with values from the `data` map.
3. **Rendering**: The `Execute` method renders the template with the data and returns the result as a string.

### Conclusion

This example demonstrates how to use Go's `html/template` package for rendering dynamic HTML content. By passing data to templates, you can easily generate HTML pages with dynamic content, ensuring a clean separation of presentation and business logic in your Go applications.
