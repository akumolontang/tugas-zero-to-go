package main

import "fmt"


func main() {
	// var x int = 13
	// var totalFaktor int = 0
	// for i:= 1;i<=x; i++ {
    // 	if x%i==0 {
	// 		totalFaktor++
	// 	   }
		   
	// }
	// if totalFaktor==2 {
	// 	fmt.Println("Bilangan Prima")
	// } else { 
	// 	fmt.Println("Bukan Bilangan Prima")}
var kata string = "namelas"
var jumlah int = len(kata)

for i:=0; i<jumlah; i++{
	if kata[i] != kata[jumlah-i-1]{
			fmt.Println("Non-Palindrome")
			break
	}else{
			fmt.Println("Palindrome")
			break
	 	}
	 } 
}
