package main

import "fmt"

type Car struct {

	Name string

	Model string

	Color string

	WeigtInKG int
}

func main () {

	c := Car{
		Name: "Ferrari",
		Model: "GTC4",
		Color: "Red",
		WeigtInKG: 1928,
	}
	fmt.Println("Car: ", c)
}