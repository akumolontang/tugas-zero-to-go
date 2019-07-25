package main

import "fmt"


func main() {
	var x int = 25
	for i:= 1;i<=x; i++ {
		for a:= 1;a<=x-i; a++{
			fmt.Print(" ")
		}
			for b:=1;b<=i; b++{
				fmt.Print(" *")
			}
	fmt.Println(" ")
}
}
	

