package main

import (
	"fmt"
	"github.com/TotallyGamerJet/w32"
)

//A simple message box example
func main() {
	time, err := w32.GetSystemTime()
	if err != nil {
		fmt.Println("Error with getting time:", err.Error())
		//return
	}
	fmt.Println(time)

}
