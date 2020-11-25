package calc

func Add (vals ...int) int {

	sum := 0

	for _, val := range vals {
		sum = sum + val
	}

	return sum
}