# goutils
Golang utils

## string

### SpliteAfter

## slice

### SliceRemove

## file

### ReadLines

    package main
    import (
        	"fmt"
        	"github.com/xuender/goutils"
    )    
    func read(line string) {
        	fmt.Print(line)
    }
    func main() {
        	goutils.ReadLines("/tmp/file.txt", read)
    }

