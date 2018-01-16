package main

import "fmt"

func wrong() {
	for i := 0; i < 10; i++ {
		defer func() {	//memory leak as defer is called inside loop
			fmt.Printf("defer runs %d\n", i)
		}()

		//defer function doesn't run here
		fmt.Printf("for step number %d\n", i)
	}
	// defer runs here
}

func solution1() {
	for i := 0; i < 10; i++ {
		func() {
			defer func() {	//defer is called after a function call, here the defer function is called after the end of the anonymous function
			//defer belongs to a func not to a block
				fmt.Printf("process(): defer runs #%d\n", i)
			}()
			fmt.Printf("processing #%d\n", i)
		}()
		fmt.Printf("for step #%d ends\n", i)
	}
}

func main() {
	wrong()
	solution1()
}
