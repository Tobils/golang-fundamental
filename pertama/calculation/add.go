package calculation

// nama func diawali huruf besar 
func Add(number int, numberTwo int) int {
	return add(number,numberTwo)
}

// hanya bisa di panggil di package yang sama
func add(number int, numberTwo int) int {
	return number + numberTwo
}