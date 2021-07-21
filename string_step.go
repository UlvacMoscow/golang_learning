package main

import "fmt"

func main()	{
	textHello := "Hello"
	textHo := "Hoho"
	i := 10
	fmt.Println(fmt.Sprintf("%s %d %s", textHello, i, textHo))
	fmt.Println(fmt.Sprintf("%T %T %T", textHello, i, textHo))
}