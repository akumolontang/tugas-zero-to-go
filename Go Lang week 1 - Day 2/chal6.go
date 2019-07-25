package main

import	"fmt"

func main() {	
		list := map[int]int{}
		list[1]=3
		list[2]=3
		list[3]=4
		list[4]=2
		list[5]=4
		list[6]=4
		list[7]=2
		list[8]=4
		list[9]=4
		// var absolute int = len(list)/2
		var freq [9]int
		
		// fmt.Println(absolute)
		for i:=1;i<=len(list);i++{
			val:= list[i]
			freq[val] = freq[val] + 1
		}
		var maxelement int = 0
		for a:=0; a<len(freq);a++{
			maxelement++
			if 	maxelement<freq[a]{
				maxelement=freq[a]
			}
		}
	fmt.Println(freq)
}