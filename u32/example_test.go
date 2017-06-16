package u32

import "fmt"

func ExampleNewSet() {
	set := NewSet(1, 2, 3, 3, 2, 1)
	fmt.Println(set.Numbers())
	// Output:
	// [1 2 3]
}

func ExampleSetCount_Count() {
	sc := SetCount{}
	sc.Add(NewSet(1, 2, 3))
	sc.Add(NewSet(2, 3, 4), NewSet(4, 5, 6))
	fmt.Println(len(sc))
	fmt.Println(sc.Count())
	// Output:
	// 6
	// [{2 2} {3 2} {4 2} {1 1} {5 1} {6 1}]
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
	// true
}

func ExampleSet_HasAll() {
	set := NewSet(6, 3)
	fmt.Println(set.HasAll(3))
	fmt.Println(set.HasAll(6))
	fmt.Println(set.HasAll(5))
	fmt.Println(set.HasAll(3, 6))
	fmt.Println(set.HasAll(3, 6, 5))
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

func ExampleSet_Jaccard() {
	a := NewSet(1, 2, 3)
	b := NewSet(1, 2, 4)
	c := NewSet(1, 4, 5)
	d := NewSet(1, 4, 5, 6, 7, 8)
	fmt.Println(a.Jaccard(b))
	fmt.Println(a.Jaccard(c))
	fmt.Println(a.Jaccard(d))
	// Output:
	// 500
	// 200
	// 125
}

func ExampleSet_Len() {
	set := NewSet(1, 2, 3, 3)
	fmt.Println(set.Len())
	// Output:
	// 3
}
