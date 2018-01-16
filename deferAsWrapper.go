package main

import "fmt"


//open a database connection, then run some queries and at the end to ensure that it gets disconnected.

type database struct {}

func (db *database) connect() (disconnect func()){
	fmt.Println("Connect")

	return func() {
		fmt.Println("disconnect")
	}
}

func wrongWay() {
	db := &database{}
	defer db.connect()
	fmt.Println("query db..")
}

func rightWay() {
	db := &database{}
	close := db.connect()
	defer close()

	fmt.Println("query db..")
}

func badPractice() {
	db := &database{}
	defer db.connect()()
	fmt.Println("query db..")
}

func main() {
	fmt.Println("wrong way\n_____________________")
	wrongWay()
	fmt.Println("Right way\n_____________________")
	rightWay()
	fmt.Println("bad Practice\n__________________")
	badPractice()
}