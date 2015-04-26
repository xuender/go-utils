package goutils

import (
	"bufio"
	"io"
	"os"
)

// Read one line, where “line” is a sequence of bytes ending with \n.
func ReadLines(file string, read func(string)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadString('\n')
		read(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
