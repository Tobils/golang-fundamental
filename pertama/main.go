package main

import (
	"fmt"
	"pertama/calculation"
)

func main(){
	fmt.Println("hallo, belajar Golang")

	sentence := TestAja()
	fmt.Println(sentence)

	result := calculation.Add(1,2)
	fmt.Println(result)
}