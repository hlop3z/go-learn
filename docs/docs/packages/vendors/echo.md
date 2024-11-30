# ECHO

Package main demonstrates a simple web server using the Echo framework.

## This server provides the following functionality

1. Serves a personalized `Hello, World!` message on the root ("/") endpoint.
2. Handles a `POST` request to create and validate user data on the "/users" endpoint.
3. Logs all registered routes to a `JSON` file.

```go
package main

import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "os"

 "github.com/go-playground/validator/v10"
 "github.com/labstack/echo/v4"
)

// User represents the structure for user data.
// It includes:
// - Name: The name of the user (required).
// - Email: The email address of the user (required and must be in a valid email format).
type User struct {
 Name  string `json:"name" validate:"required"`
 Email string `json:"email" validate:"required,email"`
}

// CustomValidator wraps the validator library to implement Echo's Validator interface.
// This allows request payloads to be validated automatically in handlers.
type CustomValidator struct {
 validator *validator.Validate
}

// Validate applies validation rules to the given struct and returns an HTTP error
// if validation fails.
func (cv *CustomValidator) Validate(i interface{}) error {
 if err := cv.validator.Struct(i); err != nil {
  // Return a 400 Bad Request error with the validation error message.
  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
 }
 return nil
}

// Get Demo
func pageHome(c echo.Context) error {
 // Get the "name" query parameter from the request.
 name := c.QueryParam("name")
 if name == "" {
  name = "World"
 }

 // Format and respond with an HTML message.
 message := fmt.Sprintf("<strong>Hello, %v!</strong>", name)
 return c.HTML(http.StatusOK, message)
}

// Post Demo
func pageUser(c echo.Context) error {
 // Bind the incoming JSON payload to the User struct.
 u := new(User)
 if err := c.Bind(u); err != nil {
  // Return a 400 Bad Request error if binding fails.
  return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
 }

 // Validate the User struct using the custom validator.
 if err := c.Validate(u); err != nil {
  // Return validation errors if the user data is invalid.
  return err
 }

 // Respond with the validated user data in JSON format.
 return c.JSON(http.StatusOK, u)
}

// initializeRoutes initializes the routes for the Echo application.
// It also logs all registered routes to a JSON file.
func initializeRoutes(app *echo.Echo) {
 // Define the root ("/") endpoint.
 app.GET("/", pageHome)

 // Define the "/users" endpoint for user data validation.
 app.POST("/users", pageUser).Name = "get-user"

 // Define the static ("/static") endpoint.
 app.Static("/static", "static")

 // Log all registered routes to a JSON file.
 routesFile := "data/routes.json"
 if err := logRoutesToFile(app, routesFile); err != nil {
  log.Printf("Failed to log routes to file: %v", err)
 }
}

// logRoutesToFile logs all registered routes to a specified JSON file.
func logRoutesToFile(app *echo.Echo, filename string) error {
 routes := app.Routes()
 data, err := json.MarshalIndent(routes, "", "  ")
 if err != nil {
  return fmt.Errorf("failed to marshal routes: %w", err)
 }

 if err := os.WriteFile(filename, data, 0644); err != nil {
  return fmt.Errorf("failed to write routes to file: %w", err)
 }

 log.Printf("Routes logged to %s", filename)
 return nil
}

// main initializes and starts the Echo web server.
func main() {
 // Create a new Echo instance.
 app := echo.New()

 // Register a custom validator.
 app.Validator = &CustomValidator{validator: validator.New()}

 // Initialize the routes.
 initializeRoutes(app)

 // Start the server on port 1323.
 app.Logger.Fatal(app.Start(":1323"))
}
```
