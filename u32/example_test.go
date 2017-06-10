package u32

import "fmt"

func ExampleNewSet() {
	set := NewSet(1, 2, 3, 3, 2, 1)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
}

func ExampleAdd() {
	set := NewSet(1, 2)
	set.Add(4)
	set.Add(3, 2, 1, 4)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3 4]
}

func ExampleClear() {
	set := NewSet(1, 2, 3)
	fmt.Println(set.Numbers())
	set.Clear()
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
	// []
}
