package utils

import "fmt"

func ExampleReadLines() {
	if err := ReadLines("LICENSE", func(line string) {
		fmt.Print(line)
	}); err != nil {
		fmt.Println(err)
	}
}

func ExampleUniqueString() {
	fmt.Println(UniqueString("U-"))
	fmt.Println(UniqueString("K-"))
	fmt.Println(UniqueString("K-"))
	// Output:
	// U-1
	// K-2
	// K-3
}

func ExampleUniqueUint32() {
	fmt.Println(UniqueUint32())
	fmt.Println(UniqueUint32())
	// Output:
	// 4
	// 5
}

func ExampleNewChMap() {
	chMap := NewChMap()
	defer chMap.Close()
	chMap.Put("key", "value")
	fmt.Println(chMap.Count())
	v, ok := chMap.Get("key")
	fmt.Println(v, ok)
	fmt.Println(chMap.Has("key"))
	fmt.Println(chMap.Has("no key"))
	// Output:
	// 1
	// value true
	// true
	// false
}

func ExampleChMap_Iterator() {
	chMap := NewChMap()
	defer chMap.Close()
	chMap.Put("key", "value")
	chMap.Iterator(func(k, v interface{}) {
		fmt.Printf("k: %s, v: %s\n", k, v)
	})
	// Output:
	// k: key, v: value
}
