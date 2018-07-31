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

// ReadBuf read file from buf.
func ReadBuf(path string, read func([]byte)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 102400)
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		read(buf[:n])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
