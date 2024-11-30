# Go HTML Template Rendering with Minification

This Go program demonstrates how to use the `html/template` package for rendering dynamic HTML pages and minifying the resulting HTML using the `github.com/tdewolff/minify` package.

The program defines an HTML template, renders it with dynamic data, and then minifies the output to reduce its size for better performance in web environments.

## Key Features

- **HTML Template Rendering**: Using Go's `html/template` package to dynamically insert values into an HTML template.
- **Minification**: After rendering the template, the output HTML is minified to remove unnecessary whitespace and optimize the content for the web.

## Example Template

### **Template (in the Go code):**

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>{{.Heading}}</h1>
    <p>{{.Content}}</p>
  </body>
</html>
```

In this template:

- `{{.Title}}`, `{{.Heading}}`, and `{{.Content}}` are placeholders for dynamic content that will be replaced at runtime.

### **Go Code Example**

Here is the Go code that reads the template, renders it with dynamic data, and minifies the final HTML output.

```go
package main

import (
 "bytes"
 "fmt"
 "html/template"
 "log"

 "github.com/tdewolff/minify/v2"
 "github.com/tdewolff/minify/v2/html"
)

// Dict is a type alias for a map with string keys and values of any type.
type Dict = map[string]interface{}
type List = []interface{}

func main() {
 // Define the template string with placeholders for dynamic content.
 var tmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Heading}}</h1>
    <p>{{.Content}}</p>
</body>
</html>
 `

 // Data structure holding dynamic values for rendering.
 data := Dict{
  "Title":   "My Go HTML Template", // Dynamic value for the title.
  "Heading": "Welcome to Go Templates!", // Dynamic value for the heading.
  "Content": template.HTML("<strong>This is a dynamic HTML page generated using the html/template package.</strong>"), // HTML content.
 }

 // Render the HTML template with the provided data.
 render := HTML(tmpl) // Returns a function to render the template.
 htmlContent := render(data) // Generate the HTML by passing the data.

 // Minify the rendered HTML content.
 fmt.Println("Minified HTML:")
 fmt.Println(minifyHTML(htmlContent))
}

// HTML takes a template string and returns a function that renders the template with the given data.
// This function allows you to reuse the template rendering logic with different data sets.
func HTML(tmpl string) func(data interface{}) string {
 // Parse the template string into a Template object.
 t, err := template.New("webpage").Parse(tmpl)
 if err != nil {
  // Panic if there's an error parsing the template.
  panic(err)
 }

 // Return a function that renders the template with the provided data.
 return func(data interface{}) string {
  var buf bytes.Buffer
  // Execute the template, passing the data to replace placeholders.
  err := t.Execute(&buf, data)
  if err != nil {
   // Panic if there's an error executing the template.
   panic(err)
  }

  // Return the generated HTML content as a string.
  return buf.String()
 }
}

// minifyHTML takes the rendered HTML string and minifies it by removing unnecessary whitespace.
func minifyHTML(htmlCode string) string {
 // Initialize the minifier.
 m := minify.New()

 // Add the minification function for HTML.
 m.AddFunc("text/html", html.Minify)

 // Minify the HTML string.
 minifiedHTML, err := m.String("text/html", htmlCode)
 if err != nil {
  log.Fatal(err) // Log and terminate if there's an error.
 }
 return minifiedHTML
}
```

## Key Components

### 1. **HTML Template Rendering**

In the `main` function, we define a template string that includes placeholders for dynamic data, such as the title, heading, and content. We then use the `HTML` function to parse and render the template with provided data.

- **Dynamic Data**: The `data` map contains values for the placeholders in the template (`Title`, `Heading`, `Content`).
- **Rendering**: The `HTML` function compiles the template and generates the final HTML string by executing the template with the given data.

### 2. **Minification**

After rendering the HTML, the `minifyHTML` function is called to minimize the output HTML. This step removes unnecessary whitespace and reduces the size of the HTML file, making it more suitable for web usage.

- **Minifier Setup**: The `github.com/tdewolff/minify` package is used for minification. Specifically, the `html.Minify` function is applied to minify the HTML.
- **Error Handling**: The function uses `log.Fatal` to terminate the program if an error occurs during minification.

### 3. **Output**

The program outputs the minified HTML to the console. For example:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>My Go HTML Template</title>
  </head>
  <body>
    <h1>Welcome to Go Templates!</h1>
    <p>
      <strong
        >This is a dynamic HTML page generated using the html/template
        package.</strong
      >
    </p>
  </body>
</html>
```

## Explanation of the Code

1. **Template Parsing**:

   - The template string is parsed using `template.New("webpage").Parse(tmpl)`.
   - If there is an error during parsing, the program panics.

2. **Template Execution**:

   - The `Execute` method is called with a buffer (`bytes.Buffer`) to capture the rendered HTML content.
   - Data is passed to the template to replace placeholders like `{{.Title}}`, `{{.Heading}}`, and `{{.Content}}`.

3. **Minification**:

   - After rendering the HTML, the `minifyHTML` function is used to reduce the size of the HTML output by removing unnecessary whitespace and line breaks.

4. **Error Handling**:
   - Both the template parsing and execution steps include error handling. If any errors occur, the program will panic (terminate).
   - The minification process uses `log.Fatal` to handle errors gracefully.

## Conclusion

This Go program demonstrates a simple but effective way to:

- Render dynamic HTML content using Go's `html/template` package.
- Minify the resulting HTML to reduce its size for better performance on the web.

You can extend this approach by adding more complex templates and handling other content types, such as CSS or JavaScript, with similar minification strategies.
