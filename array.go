package main

import "fmt"


func main()	{
	array :=[3]int{3,4,5}

	fmt.Println(len(array))
	
	for  k, i := range array {
		fmt.Println(k, i)
	}
	fmt.Println(array)
}