package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main()	{
	fmt.Println("Сколько тебе лет ?  ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	age, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	if age >= 10 {
		fmt.Println(" Age more 10")
	} else {
		fmt.Println(" Go out !!!")
	}

	fmt.Println(age)
}
