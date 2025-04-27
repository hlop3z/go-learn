package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
@BaseTypes:
- boolean
- array
- object
- integer
- number
- string
*/

type Config struct {
	Name   string  `default:"Default Name"`
	Age    int     `default:"1000000000000000000"`
	Email  string  `default:"example@domain.com"`
	Rating float64 `default:"4.5"`
}

func SetDefaults(config interface{}) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("default")

		if tag != "" && field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(tag)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				// Convert string to int64
				if intValue, err := strconv.ParseInt(tag, 10, 64); err == nil {
					field.SetInt(intValue)
				} else {
					fmt.Printf("Failed to parse int for field <%s>: %v\n", t.Field(i).Name, err)
				}
			case reflect.Float32, reflect.Float64:
				// Convert string to float64
				if floatValue, err := strconv.ParseFloat(tag, 64); err == nil {
					field.SetFloat(floatValue)
				} else {
					fmt.Printf("Failed to parse float for field <%s>: %v\n", t.Field(i).Name, err)
				}
			}
		}
	}
}

/*
 */
func testSetter() {
	cfg := &Config{}
	SetDefaults(cfg)
	fmt.Println(cfg)
}
