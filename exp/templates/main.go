package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	// Parse html file
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		fmt.Printf("Error parsing gohtml template: %v\n", err)
	}

	// Build data object to pass into html file
	data := struct {
		Name string
		Date string
		Age  int
	}{"Nathaniel Rand", "11/16/2019", 26}

	// Execute parsed file and pass in data object
	err = t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("Error executing template and data: %v\n", err)
	}
}
