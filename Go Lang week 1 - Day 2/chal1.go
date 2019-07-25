package main

import "fmt"

func main() {	
var arrayInput = [5] int {3,5,7,5,3}
var sum int

for i:=0;i<len(arrayInput);i++{
	sum = sum + arrayInput[i]
	}
var mean int = sum/len(arrayInput)
fmt.Println(mean)
}