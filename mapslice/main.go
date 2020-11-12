package main

import "fmt"

func main(){
	students := []map[string]string{
		{"name": "Agung", "Score": "A"},
		{"name": "Suhada", "Score": "A"},
		{"name": "Mario", "Score": "E"},
		{"name": "Fira", "Score": "C"},
	}

	fmt.Println(students)

	for _,student := range students {
		fmt.Println(student["name"], student["Score"])
	}
}