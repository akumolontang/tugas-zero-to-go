package main

import ( 
	"fmt"
)	


func Prime(x int) string {
	var totalFaktor int = 0
	for i:= 1;i<=x; i++ {
    	if x%i==0 {
			totalFaktor++
		   }
	}
	var primex string
	if totalFaktor==2 {
		primex = "Bilangan Prima"
	} else { 
		primex = "Bukan Bilangan Prima"
	}
		return primex
} 
func main(){
	fmt.Println(Prime(1))
	
	fmt.Println(Prime(3))

	fmt.Println(Prime(7))

	fmt.Println(Prime(6))

	fmt.Println(Prime(23))

}