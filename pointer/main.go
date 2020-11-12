package main

import "fmt"

func main() {
	numberA := 5
	numberB := &numberA // menyimpan alamat memory reference
	numberC := *numberB // dereference

	fmt.Println(numberB)
	fmt.Println(numberC)
	fmt.Println(numberA)

	// merubah nilai di alamat
	*numberB = 10
	fmt.Println(*numberB)
	fmt.Println(numberA)

	var numA int = 5
	var numB *int = &numA

	fmt.Println(*numB)
	fmt.Println(numA)

	numA = 20

	fmt.Println(*numB)
	fmt.Println(numA)

	num1 := 5
	fmt.Println("alamat memory ", &num1)
	fmt.Println("nilai awal ", num1)
	change(&num1, 100)
	fmt.Println("nilai akhir ", num1)
	fmt.Println("alamat memory ", &num1)

}

func change(old *int, new int) {
	*old = new
	fmt.Println("alamat memory ", old)
	// fmt.Println("nilali di dalam fungsi ", old)
}
