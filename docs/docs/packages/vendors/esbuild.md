# ESBuild (**TSX**)

```go
package main

import (
 "fmt"
 "regexp"
 "strings"

 "github.com/evanw/esbuild/pkg/api"
)

func main() {
 // Sample TSX input
 tsxCode := `
import "./style.scss";

// @GO

const Layout = (props) => (
  <div>
    <header>Head</header>
    <main>{props.children}</main>
    <footer>Head</footer>
  </div>
);

export default function Button(props) {
 const regexNoReplace = "export default"
  return (
    <Layout x-html {...props} class={[$NAME, props.class]}>
      {props.children}
    </Layout>
  );
}

    `

 customizedCode := componentJSX("Component", tsxCode)
 fmt.Println("Transformed Code:")
 fmt.Println(customizedCode)
}

func parseTSX(input string, methodName string, minify bool) string {
 if methodName == "" {
  methodName = "React.createElement" // Default value
 }

 // Transform TSX to JavaScript
 result := api.Transform(input, api.TransformOptions{
  Loader: api.LoaderTSX,    // Set the loader to TSX
  JSX:    api.JSXTransform, // Enable JSX transformation
  Format: api.FormatESModule,
 })

 if len(result.Errors) > 0 {
  fmt.Println("Errors:", result.Errors)
  return ""
 }

 // Replace `/* @__PURE__ */ React.createElement` with a custom function
 transformedCode := string(result.Code)
 customizedCode := strings.ReplaceAll(
  transformedCode,
  "/* @__PURE__ */ React.createElement",
  methodName,
 )

 if minify {
  return minifyJS(customizedCode)
 }

 return customizedCode
}

// componentJSX wraps a TSX transformation to create a reusable component function in JavaScript.
// It generates a JavaScript function named after the provided `name` parameter and applies
// the transformation and minification logic.
//
// Parameters:
// - name: The name of the JavaScript function to be generated.
// - input: The TSX input code to be processed.
//
// Returns:
// - A minified JavaScript function string representing the transformed component.
func componentJSX(name string, input string) string {
 // Fix Code Pre-String
 tsxCode := strings.Split(input, "// @GO")[1]
 tsxCode = strings.TrimSpace(tsxCode)

 // Call the function to replace `export default` only at the start of a line
 tsxCode = replaceRegex(tsxCode, `(?m)^export default`, "return")
 // Wrap the transformed code into a named component function
 tsxCode = fmt.Sprintf("const %s = (function() { %s })();", name, tsxCode)

 // Transform the input TSX and use "h" as the custom JSX pragma
 transformedCode := parseTSX(tsxCode, "h", false)

 return minifyJS(transformedCode)
}

// minifyJS reduces the size of the given JavaScript code by applying whitespace,
// identifier, and syntax minification using esbuild.
//
// Parameters:
// - input: The JavaScript code to be minified.
//
// Returns:
// - A minified JavaScript string, or an empty string if an error occurred.
func minifyJS(input string) string {
 result := api.Transform(input, api.TransformOptions{
  MinifyWhitespace:  true,
  MinifyIdentifiers: true,
  MinifySyntax:      true,
 })

 if len(result.Errors) > 0 {
  fmt.Println("Errors:", result.Errors)
  return ""
 }

 return string(result.Code)
}

// replaceExportDefault replaces occurrences of "export default" with the given replacement
// string, but only if "export default" is at the beginning of a line.
//
// Parameters:
// - input: The input string containing code.
// - replacement: The string to replace "export default" with.
//
// Returns:
// - The updated string with replacements applied.
func replaceRegex(input, query string, replacement string) string {
 // Regex pattern to match "export default" at the start of a line
 // `(?m)^export default`:
 //   - `(?m)`: Enables multiline mode to allow `^` to match the start of each line.
 //   - `^export default`: Matches "export default" at the beginning of a line.
 re := regexp.MustCompile(query)

 // Perform the replacement
 return re.ReplaceAllString(input, replacement)
}
```
