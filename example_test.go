package goutils

import "fmt"

func ExampleReadLines() {
	ReadLines("LICENSE", func(line string) {
		fmt.Print(line)
	})
}

func ExampleSplitAfter() {
	fmt.Println(SplitAfter("110120119129", "0", "9"))
	// Output:
	// [110 120 119 129]
}
