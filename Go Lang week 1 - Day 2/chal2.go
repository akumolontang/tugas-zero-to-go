package main

import "fmt"

func main() {	
var sliceInput = []float32 {1,2,3,4,5}
var sum float32
var numberslice int = len(sliceInput)/2
var median float32
if len(sliceInput)%2==0{
	sum = sum + sliceInput[numberslice] + sliceInput[numberslice-1]
	median = sum/2
	}else{
	sum = sliceInput[numberslice]
	}
	fmt.Println(median)
}