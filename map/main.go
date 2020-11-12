package main

import "fmt"

func main(){
	/**
	key => string
	val => int
	*/
	var myMap map[string]int
	myMap = map[string]int{} // define var
	myMap["ruby"] = 9
	myMap["go"] = 10
	myMap["javascript"] = 8
	myMap["java"] = 10
	fmt.Println(myMap)
	fmt.Println(myMap["ruby"])


	/**
	deinfe 2
	*/
	secondMap := map[string]string{"ruby":"is beautiful", "go":"is super fast"}
	fmt.Println(secondMap)

	for key, value := range secondMap {
		fmt.Println("key :", key, "value :", value)
	}

	delete(secondMap, "go")

	for key, value := range secondMap {
		fmt.Println("key :", key, "value :", value)
	}


	
}