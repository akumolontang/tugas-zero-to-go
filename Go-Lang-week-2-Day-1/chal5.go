package main

import ( 
	"fmt"
)	
func Geom(arr []int) bool {
	var jumlah int = len(arr)-1
	var check1 = arr[1] / arr[0]
	var boll bool
	for i:=1;i<jumlah;i++{
	if arr[i]/arr[i-1] == check1 {
				boll = true
		}else{
				boll = false
				break
		}
	}
	return boll
} 
		

func main(){
	fmt.Println(Geom([]int{1,3,9,27,81}))
	
	fmt.Println(Geom([]int{2,4,8,16,32}))
	
	fmt.Println(Geom([]int{2,4,7,8}))

	fmt.Println(Geom([]int{2,6,18,54}))

	fmt.Println(Geom([]int{1,2,3,4,7,9}))
}