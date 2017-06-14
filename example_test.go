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

func ExamplePrefixBytes() {
	fmt.Println(PrefixBytes("id-", []byte("abc")))
	// Output:
	// [105 100 45 97 98 99]
}

func ExamplePrefixUint32() {
	fmt.Println(PrefixUint32("id-", 3))
	// Output:
	// [105 100 45 3 0 0 0]
}

func ExampleUniqueId() {
	fmt.Println(UniqueId("U-"))
	fmt.Println(UniqueId("K-"))
	fmt.Println(UniqueId("K-"))
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

	chMap.Set("key", "value")
	v, ok := chMap.Get("key")
	fmt.Println(v, ok)
	fmt.Println(chMap.Has("key"))
	fmt.Println(chMap.Has("no key"))
	// Output:
	// value true
	// true
	// false
}
