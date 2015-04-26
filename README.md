# goutils
Golang utils

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


