//go:build windows
// +build windows

package main

import "github.com/muesli/termenv"

func main() {
	mode, err := termenv.EnableWindowsANSIConsole()
	if err != nil {
		panic(err)
	}
	defer termenv.RestoreWindowsConsole(mode)

	menuLoop()
}
