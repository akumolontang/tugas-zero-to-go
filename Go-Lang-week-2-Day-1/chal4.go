package main

import ( 
	"fmt"
)	

func Palindrome(s string) bool {
	var jumlah int = len(s)
	var boll bool
	for i:=0; i<jumlah; i++{
		if s[i] != s[jumlah-i-1]{
				boll = false
				break
		}else{
				boll = true
				break
			 }
		 } 
		 return boll
	}
	

func main(){
	fmt.Println(Palindrome("katak"))
	fmt.Println(Palindrome("blanket"))
	fmt.Println(Palindrome("civic"))
	fmt.Println(Palindrome("kasur rusak"))
	fmt.Println(Palindrome("mister"))


}