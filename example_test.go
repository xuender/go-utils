package utils

import "fmt"

func ExampleReadLines() {
	if err := ReadLines("LICENSE", func(line string) {
		fmt.Print(line)
	}); err != nil {
		fmt.Println(err)
	}
}

func ExampleSplitAfter() {
	fmt.Println(SplitAfter("110120119129", "0", "9"))
	// Output:
	// [110 120 119 129]
}

func ExamplePrefixBytes() {
	fmt.Println(PrefixBytes([]byte("id-"), []byte("abc")))
	// Output:
	// [105 100 45 97 98 99]
}

func ExamplePrefixUint32() {
	fmt.Println(PrefixUint32([]byte("id-"), 3))
	// Output:
	// [105 100 45 3 0 0 0]
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

func ExampleJoin() {
	a := []byte("aa")
	b := []byte("bb")
	fmt.Println(Join([]byte(","), a, b))
	// Output:
	// [97 97 44 98 98]
}

func ExampleConcat() {
	a := []byte("aa")
	b := []byte("bb")
	fmt.Println(Concat(a, b))
	// Output:
	// [97 97 98 98]
}
