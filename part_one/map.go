package main

import "fmt"

func main ()	{

	cart := map[int]int{
		1: 1,
		2: 4,
	}

	for key, value := range cart {
		fmt.Println(key, value)
	}
}