package main

import (
	"github.com/totallygamerjet/w32"
)

//A simple message box example
func main() {
	w32.MessageBox(0, "Hello World!", "Hello, World!", 0)
}
