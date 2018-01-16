package main

import "fmt"

type Car struct {
	model	string
}

func (car Car) PrintModel() {
	fmt.Println("car.Models == ", car.model)
}

func (car *Car) PrintPointerModel() {
	fmt.Println("car.Model == ", car.model)
}

func withoutPointer() {
	c := Car{model:"Ferrari"}
	defer c.PrintModel()
	c.model = "BMW"
}

func withPointer() {
	c := Car{model:"Ferrari"}
	defer c.PrintPointerModel()
	c.model = "BMW"
}

func main() {
	withoutPointer()
	withPointer()
}
