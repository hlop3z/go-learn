package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Dict = map[string]interface{}

// Blueprint struct to store dynamic methods and other data
type Blueprint struct {
	methods map[string]func(args Dict) interface{}
}

// Constructor to initialize a new Blueprint
func NewBlueprint() *Blueprint {
	return &Blueprint{
		methods: make(map[string]func(args Dict) interface{}),
	}
}

func cleanName(input string) string {
	// Compile a regular expression to match characters that are not letters, numbers, or underscores
	re := regexp.MustCompile(`[^\w]`)
	// Replace non-matching characters with an empty string
	cleaned := re.ReplaceAllString(input, "")
	// Convert the cleaned string to lowercase
	return strings.ToLower(cleaned)
}

// AddMethod dynamically adds a method to the Blueprint
func (bp *Blueprint) AddMethod(name string, method func(args Dict) interface{}) {
	bp.methods[cleanName(name)] = method
}

// CallMethod invokes a dynamic method by name
func (bp *Blueprint) CallMethod(name string, args Dict) (interface{}, error) {
	// Success
	if fn, exists := bp.methods[name]; exists {
		return fn(args), nil
	}
	// Error
	return nil, fmt.Errorf("method %s not found", name)
}

// Global Variable
var Schema = Dict{"id": 1}

func main() {
	// Initialize Blueprint
	blueprint := NewBlueprint()

	// Add methods
	blueprint.AddMethod("create", CreateMethod)
	blueprint.AddMethod("remove", RemoveMethod)
	blueprint.AddMethod("update", UpdateMethod)
	blueprint.AddMethod("delete", DeleteMethod)
	blueprint.AddMethod("detail", DetailMethod)
	blueprint.AddMethod("search", SearchMethod)

	// Call methods dynamically
	fmt.Println(Schema)

	// Call methods dynamically
	if result, err := blueprint.CallMethod("create", Dict{"key": "value"}); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	if result, err := blueprint.CallMethod("search", Dict{"query": "blueprints"}); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	// Attempt to call a non-existent method
	if result, err := blueprint.CallMethod("nonexistent", nil); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

// Predefined methods for the Blueprint
func CreateMethod(args Dict) interface{} {
	return fmt.Sprintf("Create called with args: %v", args)
}

func UpdateMethod(args Dict) interface{} {
	return fmt.Sprintf("Update called with args: %v", args)
}

func DeleteMethod(args Dict) interface{} {
	return fmt.Sprintf("Delete called with args: %v", args)
}

func RemoveMethod(args Dict) interface{} {
	return fmt.Sprintf("Remove called with args: %v", args)
}

func DetailMethod(args Dict) interface{} {
	return fmt.Sprintf("Detail called with args: %v", args)
}

func SearchMethod(args Dict) interface{} {
	return fmt.Sprintf("Search called with args: %v", args)
}
