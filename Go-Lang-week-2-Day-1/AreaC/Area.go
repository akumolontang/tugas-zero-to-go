package CALC

func Calcx(l float32, w float32) (float32,float32){
	area:= l*w
	perim:= (l+w)*2
	return area, perim
}