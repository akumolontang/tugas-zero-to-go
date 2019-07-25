package main

import "fmt"


func faktorBilangan(x int) {
	for i:= 1;i<=x; i++ {
    	if x%i==0 {
			fmt.Print(i," ")
		}
	}
	fmt.Println()
}	

func main(){
	faktorBilangan(6)
	
	faktorBilangan(12)
	
	faktorBilangan(14)
	
	faktorBilangan(16)
	
	faktorBilangan(20)
}