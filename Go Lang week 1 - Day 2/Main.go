package main

import "fmt"

// __________________ARRAY_____________
// Array is a data structure

func main() {
		// var num = 5
		// var list [5] string
//Indexing 4 -> 0 to 3
//Indexing starts from 0

		// list [0] = "hai"
		// list [1] = "kamu"
		// list [2] = "lagi"
		// list [3] = "jelek"
		// fmt.Println(list[2])

		
		// list [0] = "hai"
		// list [1] = "kamu"
		// list [2] = "lagi"
		// list [3] = "jelek"
		
		// var jumlah int = len(list)
		// fmt.Println(jumlah)
		
		// underscore is not used and is a garbage collector


		// list [0] = "hai"
		// list [1] = "kamu"
		// list [2] = "lagi"
		// list [3] = "jelek"
		// list [4] = "banget"
		
		
		// var jumlah int = len(list)
		// for i:=0; i<jumlah ;i++{
		// 	fmt.Println(list[i])	
		// }
		// list [0] = "hai"
		// list [1] = "kamu"
		// list [2] = "lagi"
		// list [3] = "jelek"
		// list [4] = "banget"
		
		
		// // var jumlah int = len(list)
		// for _, nama := range list{
		// 	fmt.Println("ke", len(list), nama)	
		// }

		// list [0] = "hai"
		// list [1] = "kamu"
		// list [2] = "lagi"
		// list [3] = "jelek"
		// list [4] = "banget"
		
		
		// var jumlah int = len(list)
		// var list =[...]string {"","1","2","3"}
		// 	fmt.Println(len(list))	
			
		// var list =[4]string {"hahaha","1","2","3"}
		// fmt.Println(list)	

        //ASCII number when printing array of string[x] 
		// var list string = "asdfgh"
		// fmt.Println(list[2])	

		// Strings return the letter from the array
		// var list string = "asdfgh"
		// fmt.Println(string(list[2]))	

		
		// var list =[2][5] int {{5,12,23,1,1},{12,24,1,25,5}}
		// fmt.Println(list[])

		
		// var list =[2][5] int {{5,12,23,1,1},{12,24,1,25,5}}
		// fmt.Println(list[1][3])


		// var list =[3][5][1] int {{{5,12,23,1,1},{12,24,1,25,5}}}
		// fmt.Println(list)

// _______________SLICE____________

		// var list []string
		// list = append(list, "hai")
		// list = append(list, "kamu")
		// list = append(list, "lagi")
		// list = append(list, "jelek")

		// var list []string
		// list = append(list, "hai","kamu","jancuk")


		// list:= []string {"bangsat","kamu"}
		// var list2 = []string{"bangsat","kau"}
		// fmt.Println(list,list2)
		
		// list:= []string {"bangsat","kamu"}
		// var list2 = []string{"bangsat","kau","jelek"}
		// fmt.Println(len(list),len(list2))
		
		// var list = []string{"saya","minum","susu"}
		// list2 := list[0:3]
		// fmt.Println(list2)
		// fmt.Println(list)
		
		// 
		// var list = []string{"saya","minum","susu"}
		// list2 := list[0:2]
		// list2[1] = "keode" //changes the parent slice//
		// fmt.Println(list2)
		// fmt.Println(list)


		// var list = []string{"saya","minum","susu"}
	
		// for i:=0;i<len(list); i++{
		// 	fmt.Println(list[i])
		// }


		// var list = []string{"saya","minum","susu"}
	
		//index no need to declare but is 0-number
		//name declared range is all the array
		// for index, name:= range list{
		// 	fmt.Println(index+1,name)
		// }
		// ____CAVEAT:
		// name can only be used inside 'for'

		// var list = []string{"saya","minum","susu"}	
		// for index, name:= range list{
		// 	fmt.Println(index+1,name)
		// }

		// var list = []string{"saya","minum","susu"}	
		// var index int
		// var name string
		// for index, name:= range list{
		// 	fmt.Println(index+1,name)
		// }
		// fmt.Println(index,name)

		// var list = []string{"saya","minum","susu","sapi"}	
		// list2 := list[:2]
		// 	fmt.Println(list2)

		// var list = []string{"saya","minum","susu","sapi"}	
		// list2 := list[0:4]
		// 	fmt.Println(list2)

		// var list = []string{"Time","minum","susu","sapi"}	
		// list2 := list[0:4]
		// 	fmt.Println(list2)

		// var list = []string{"Time","minum","susu","sapi"}	
		// for i:=0;i<len(list),i++{
		// 	fmt.Println(list[1])
		// }

		
		// var list = []string{"Time","minum","susu","sapi"}	
		// for index, name := range list{
		// 	fmt.Println(index,name)
		// }

			
		// var list = []string{"same","minum","susu","sapi"}	
		// for _, name := range list{
		// 	fmt.Println(name)
		// }
		// fmt.Println(list[2:4])

		// var list = []string{"same","minum","susu","sapi"}	
		// for _, name := range list{
		// 	fmt.Println(name)
		// }
		// fmt.Println(list[2:4])

// ___________MAP_______________
//Unique value defined -- value inside bracket cannot be the same
		// var list = make(map[string]string)
		// list["hai"]	
		
		// var list = make(map[string]string)
		// list["name"]="Jancuk"
		// fmt.Println(list)
		

		// list := map[string]string{}
		// list["name"]="Jancuk"
		// fmt.Println(list)

	
		// var list map[string]string
		// list = map[string]string{}
		// list["name"]="Jancuk"

		// list := map[bool]string{}
		// list[true]="MANTUL BOSKU"
		// list[false]="BANGSAT KAU"
		// fmt.Println(list)
		// if	2==3{
		// 	fmt.Println(list[false])
		// }else{
		// 	fmt.Println(list[true])
		// }

		// list := map[string]string{}
		// list["name"]="Jancuk"
		// list["address"]="Sebelah Warung"
		// 	for key, value:= range list{
		// 		fmt.Println(key,value)
		// 	}

		// list := map[int]string{}
		// list[0]="Jancuk"
		// list[1]="Sebelah Warung"
		// for i:=0;i<len(list);i++{
		// 	fmt.Println(list)
		// }
		
		list := map[int]string{}
		list[0]="Jancuk"
		list[1]="Sebelah Warung"
		for i:= 0;i<len(list);i++{
			fmt.Println("# ",i," ",list[i])
		}
}
