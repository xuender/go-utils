package goutils

import (
	"fmt"
	"os/exec"
	"runtime"
)

var commands = map[string]func(uri string) error{
	"windows": win,
	"darwin":  def("open"),
	"linux":   def("xdg-open"),
}

func win(uri string) error {
	cmd := exec.Command("cmd", "/c", "start", uri)
	return cmd.Start()
}
func def(cmd string) func(uri string) error {
	return func(uri string) error {
		cmd := exec.Command(cmd, uri)
		return cmd.Start()
	}
}

// Open calls the OS default program for uri
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("未知的操作系统 %s .", runtime.GOOS)
	}
	return run(uri)
}
