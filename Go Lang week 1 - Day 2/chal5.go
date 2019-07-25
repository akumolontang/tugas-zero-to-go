package main

import	"fmt"

func main() {	
		var freq [15]int
		// var mod int
		var val int
		list := map[int]int{}
		list[1]=1
		list[2]=2
		list[3]=3
		list[4]=4
		list[5]=4
		list[6]=4
		// var absolute int = len(list)/2
		
		// // fmt.Println(absolute)
		for i:=1;i<=len(list);i++{
			val = list[i]
			freq[val] = freq[val] + 1
			// fmt.Println(freq[val])
			
		
		}
		var maxelement int = 0
		for a:=0; a<len(freq);a++{
			maxelement++
			if 	maxelement>=freq[a]{
				maxelement=freq[a]
			}
		fmt.Println(list[maxelement])
			
		} 
		
			
		

}