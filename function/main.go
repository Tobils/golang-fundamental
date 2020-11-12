package main

import "fmt"

func main(){
	sentences := printMine("suhada")
	fmt.Println(sentences)

	val := add(3, 5)
	fmt.Println(val)

	luas, keliling := calculatenew(39, 12)
	fmt.Println(luas, keliling)
}

func printMine(sentence string) string {
	res := sentence + " mantaf uhuy"
	return res
}


func add(val1 int, val2 int) int {
	return val1 + val2
}


// multiple return 
func calculate(p int, l int)(int, int){
	luas := p * l
	keliling := 2 * (p + l)

	return luas, keliling
}

// with predfined return 
func calculatenew(p int, l int)(luas int,keliling int){
	luas = p * l
	keliling = 2 * (p + l)

	return
}
