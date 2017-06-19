package goutils

import (
	"bufio"
	"io"
	"os"
)

// ReadLines read file.
func ReadLines(file string, read func(string)) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() {
		e := f.Close()
		if err == nil {
			err = e
		}
	}()
	bfRd := bufio.NewReader(f)
	for {
		var line string
		line, err = bfRd.ReadString('\n')
		read(line)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
	}
}
