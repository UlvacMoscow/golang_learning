package main

import "fmt"


func main()	{
	var x[5]int
	x[4] = 100
	x[3] = 2.0
	x[0] = 5
	fmt.Println(x)
	slice_example()
	slice_example_yet()
}

func slice_example()	{
	var x[]float64
	y := make([]float64, 5)
	z := make([]float64, 5, 10)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func slice_example_yet()	{
	var arr[]float64
	arr := [5]float64{1,2,3,4,5}
}
