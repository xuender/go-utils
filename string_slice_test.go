package utils

import (
	"fmt"
)

func ExampleStringSlice_Delete() {
	ret := StringSlice([]string{"123", "234", "234"})

	fmt.Println(ret.Delete("345"))
	fmt.Println(ret.Delete("234"))
	fmt.Println(ret)

	// Output:
	// [123 234 234]
	// [123]
	// [123]
}

func ExampleStringSlice_Add() {
	ret := StringSlice([]string{"123"})

	fmt.Println(ret.Add("234", "345"))
	fmt.Println(ret)

	// Output:
	// [123 234 345]
	// [123 234 345]
}

func ExampleStringSlice_IndexOf() {
	ret := StringSlice([]string{"123", "234"})

	fmt.Println(ret.IndexOf("234"))
	fmt.Println(ret.IndexOf("345"))

	// Output:
	// 1
	// -1
}

func ExampleStringSlice_Contains() {
	ret := StringSlice([]string{"123", "234"})

	fmt.Println(ret.Contains("234"))
	fmt.Println(ret.Contains("345"))

	// Output:
	// true
	// false
}
