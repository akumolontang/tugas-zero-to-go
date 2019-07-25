package main

import ("fmt"
"Go-Lang-week 2-Day-1/CALC"
)
// ________________FUNCTION__________________________
// Function w parameter -- in this case the funct greeting(hour int) 
// is the function
// Doesn't have to be in order - say function first then execution
// Could be either first or last

// func greeting(hour int){

// if hour <12 {
// 	fmt.Println("Morning ", hour)
// }else if hour < 18 {
// 	fmt.Println("Afternoon ", hour)
// }else{
// 	fmt.Println("Night ", hour)
// }
// }

// func time(day int, name string){
// if day < 5{
// 	fmt.Println("Weekday ", name)
// }else{
// 	fmt.Println("Weekend ", name)
// }

// }

// func main() {	
// hour:=18
// greeting(hour)
// day:= 7
// name:= "Sunday"
// time(day, name)



// }




// func areax(l int, w int) float64{
// 	area:= float64(l * w)/2
// 	// var perim float64 = float64((l*2)+(w*2)) + area
// 	return area
// 	// return perim
// }
	
// 	func main() {	

// 	fmt.Println(areax(1,22))
	
// }

// ____________________________________________________________

	
	func main() {	
		l := float32(9)
		w := float32(8.11)
	resulta, resultp := CALC.Calcx(l,w)
	fmt.Println(resulta)
	fmt.Println(resultp)
}