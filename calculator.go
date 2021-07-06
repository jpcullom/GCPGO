package main

import "strconv"

func Add(x string, y string) int{
	xInt, err := strconv.Atoi(x)
	yInt, err2 := strconv.Atoi(y)

	//Error Checking
	if (err != nil) && (err2 != nil){
		panic("Both inputs cannot be converted. Idiot.")
	} else if err != nil {
		panic("First input cannot be converted. Moron.")
	} else if err2 != nil {
		panic("Second input cannot be converted. Rube.")
	}
	//End Error Checking

	return xInt + yInt
}
func DoubleTrain(slice []int) []int{
	for i := 0; i < len(slice); i++{
		temp := slice[i]*2
		slice[i] = temp
	}

	return slice
}
