package u32

import "fmt"

func ExampleNewSet() {
	set := NewSet(1, 2, 3, 3, 2, 1)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
}

func ExampleSet_Add() {
	set := NewSet(1, 2)
	set.Add(4)
	set.Add(3, 2, 1, 4)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3 4]
}

func ExampleSet_Clear() {
	set := NewSet(1, 2, 3)
	fmt.Println(set.Numbers())
	set.Clear()
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
	// []
}

func ExampleSet_Has() {
	set := NewSet(6, 3)
	fmt.Println(set.Has(3))
	fmt.Println(set.Has(6))
	fmt.Println(set.Has(5))
	fmt.Println(set.Has(3, 6))
	fmt.Println(set.Has(3, 6, 5))
	// Output:
	// true
	// true
	// false
	// true
	// false
}

func ExampleSet_Hit() {
	set := NewSet(1, 2, 3)
	fmt.Println(set.Hit(2, 3, 4, 5))
	// Output:
	// 2
}
