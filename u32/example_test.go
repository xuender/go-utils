package u32

import "fmt"

func ExampleNew() {
	set := New(1, 2, 3, 3, 2, 1)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
}

func ExampleAdd() {
	set := New(1, 2)
	set.Add(4)
	set.Add(3, 2, 1, 4)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3 4]
}

func ExampleClear() {
	set := New(1, 2, 3)
	fmt.Println(set.Numbers())
	set.Clear()
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
	// []
}
