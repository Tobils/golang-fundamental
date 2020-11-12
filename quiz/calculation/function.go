package calculation

import "errors"

func Sum(data []int) (sum int) {
	for _, val := range data {
		sum += val
	}

	return
}

// nilai default dari error adalah nil 
func Calculate(val1 int, val2 int, keyword string) (result int,errorResult error){
	switch keyword {
		case "+":
			result = val1 + val2
		case "-":
			result = val1 - val2
		case "*":
			result = val1 * val2
		case "/":
			result = val1 / val2
		default :
			errorResult = errors.New("Unknown Operation")
	}
	return
}
