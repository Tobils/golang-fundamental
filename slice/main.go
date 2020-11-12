package main


import "fmt"
/**
slice mirip dengan array tapi dak butuh panjang element
*/
func main() {

	// define 1
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "Plastatsion 4")
	gamingConsole = append(gamingConsole, "Plastatsion 2")
	gamingConsole = append(gamingConsole, "nintendo")
	fmt.Println(gamingConsole)


	// define 2
	gamers := []string{"ps", "ml", "pubg"}
	fmt.Println(gamers)
	

	for _, val := range gamers {
		fmt.Println(val)
	}

}