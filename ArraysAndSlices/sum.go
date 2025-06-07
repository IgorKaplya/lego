package main

func Sum(numbers []int) int {
	var result = 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbers ...[]int) []int {
	result := make([]int, len(numbers))

	for i, item := range numbers {
		result[i] = Sum(item)
	}

	return result
}

func SumAllTails(numbers ...[]int) []int {
	result := make([]int, len(numbers))

	for i, v := range numbers {
		result[i] = 0
		if len(v) > 0 {
			result[i] = Sum(v[1:])
		}

	}

	return result
}
