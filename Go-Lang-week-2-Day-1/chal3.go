package main

import ( 
	"fmt"
)	

func generatePattern(x int) {
	for i:= 1;i<=x; i++ {
		if i!=1{
			fmt.Println()
		}
		for a:=1;a<=i;a++{
		fmt.Print(a, " ")
		}
		}
	fmt.Println(" ")
}

func main(){
	generatePattern(5)
	

}