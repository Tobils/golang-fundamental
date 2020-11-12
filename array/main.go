package main

import (
	"fmt"
)

func main(){
	// definisi 1 
	var bahasa [5]string
	bahasa[0] = "man jadda wa jada"
	bahasa[1] = "ok"
	bahasa[2] = "ntafs"
	bahasa[3] = "gaskeun"
	bahasa[4] = "bungkus"

	fmt.Println(bahasa)
	length := len(bahasa)
	fmt.Println(length)

	for i := 0; i < len(bahasa); i++ {
		fmt.Println(bahasa[i])
	}



	// definisi 2
	languages := [5]string{"Ruby", "Python", "JavaScript", "Go", "C#"}

	for _, value := range languages {
		fmt.Println(value)
	}
}