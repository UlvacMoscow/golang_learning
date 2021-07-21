package main

import "fmt"


func main()	{
	i := 1
	for i < 10	{
		fmt.Println(i, "a")
		i += 1
	}
	other_syntax()

}

func other_syntax()	{
	for k := 1; k <=10; k++	{
		if k % 2 == 0 {
			fmt.Print("четное ")
		} else {
			fmt.Print("Нечетное ")
		}
		fmt.Println(k)
	}

}