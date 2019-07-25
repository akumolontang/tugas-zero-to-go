package main

import "fmt"

func main() {	
var sliceInput = []float32 {1,2,3,4,5}
var sum float32
var numberslice int = len(sliceInput)
var median float32
if numberslice%2==0{
	sum = sum + sliceInput[numberslice/2] + sliceInput[(numberslice/2)-1]
	median = sum/2
	}else{
	sum = sliceInput[numberslice/2]
	}
	fmt.Println(median)
}