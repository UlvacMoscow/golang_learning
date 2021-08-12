package main

import "fmt"

func main()	{
	slice := make([]int, 3)
	fmt.Println(len(slice))

	slice = append(slice, 321)
	fmt.Println(len(slice))

	for _, i := range slice {
		fmt.Println(i)
	}
	fmt.Println(slice)
}
 