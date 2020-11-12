package main

import (
	"fmt"
	"quiz/calculation"
)

func main() {
	fmt.Println("Go Lang Fundamental 1st Quiz")
	result := calculation.Multiplication(45, 67)
	fmt.Println(result)

	number := 11
	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Not Found !")
	}

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	k := 0
	for k <= 100 {
		fmt.Println(k)
		k++
	}

	kata := "Golang the best language   "
	for index, letter := range kata {
		if index%2 == 0 {
			fmt.Println(string(letter))
		}
	}

	for _, letter := range kata {
		switch letter {
		case 'a':
			fmt.Println(string('a'))
		case 'i':
			fmt.Println(string('i'))
		case 'u':
			fmt.Println(string('u'))
		case 'e':
			fmt.Println(string('e'))
		case 'o':
			fmt.Println(string('o'))
		}
	}

	// hitung rata-rata dan nilai >= 90
	scores := [9]int{100, 80, 75, 92, 70, 93, 12, 67, 15}
	tmp := 0
	for _, val := range scores {
		tmp += val
	}
	fmt.Println(tmp)
	average := float64(tmp / len(scores))
	fmt.Println(average)

	avg := calculation.Average(scores, len(scores))
	fmt.Printf("the value is :%f", avg)
	fmt.Println()

	var goodScores []int

	for _, val := range scores {
		if val >= 90 {
			goodScores = append(goodScores, val)
		}
	}

	fmt.Println(goodScores)

	nilai := []int{10, 5, 8, 9, 7}
	total := calculation.Sum(nilai)

	fmt.Println(total)

	// result, err := calculation.Calculate(10,2, "+")
	// result, err := calculation.Calculate(10,2, "-")
	// result, err := calculation.Calculate(10,2, "*")
	result, err := calculation.Calculate(10, 2, "/")
	// result, err := calculation.Calculate(10,2, "d")

	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}
