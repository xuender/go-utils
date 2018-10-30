package utils

import (
	"fmt"
)

func ExampleStringSlice() {
	ss := StringSlice([]string{"123", "234", "234"})

	fmt.Println(ss.Delete("234"))
	fmt.Println(ss)

	// Output:
	// [123]
	// [123]
}
func ExampleStringSlice_Delete() {
	ss := StringSlice([]string{"123", "234", "234"})

	fmt.Println(ss.Delete("345"))
	fmt.Println(ss.Delete("234"))
	fmt.Println(ss)

	// Output:
	// [123 234 234]
	// [123]
	// [123]
}

func ExampleStringSlice_Add() {
	ss := StringSlice([]string{"123"})

	fmt.Println(ss.Add("234", "345"))
	fmt.Println(ss)

	// Output:
	// [123 234 345]
	// [123 234 345]
}

func ExampleStringSlice_IndexOf() {
	ss := StringSlice([]string{"123", "234"})

	fmt.Println(ss.IndexOf("234"))
	fmt.Println(ss.IndexOf("345"))

	// Output:
	// 1
	// -1
}

func ExampleStringSlice_Contains() {
	ss := StringSlice([]string{"123", "234"})

	fmt.Println(ss.Contains("234"))
	fmt.Println(ss.Contains("345"))

	// Output:
	// true
	// false
}
