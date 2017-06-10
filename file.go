package goutils

import (
	"bufio"
	"io"
	"os"
)

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
