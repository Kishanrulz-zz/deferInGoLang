package main

import "fmt"

func main() {
	var run func() = nil
	defer run()		//registers a call to the stack

	fmt.Println("runs")

	//panics here because of nil function
}
