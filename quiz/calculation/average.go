package calculation

func Average(data [9]int, size int) float64 {
	var i int;
	var avg float64
	var sum int = 0

	for i=0; i<len(data); i++ {
		sum += data[i]
	}

	avg = float64(sum / size)
	return avg
}