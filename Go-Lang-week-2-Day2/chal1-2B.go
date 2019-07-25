package main

import "fmt"

type Student struct {

	RollNumber int
	Name string
}



func main () {
	s := Student{
		RollNumber: 11,
		Name: "Jack",
	}
	ps := &s
	
	fmt.Println(ps)
	fmt.Println(ps.Name)
	s.RollNumber = 31
	fmt.Println(ps.Name)
	fmt.Println(ps)
}