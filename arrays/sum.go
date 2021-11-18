package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return result
}

// Note that the size of the array is encoded in its type
func SumArray(numbers [5]int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return result
}

// This is a variadic function (https://gobyexample.com/variadic-functions)
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Slice the slice to get the last item (slice[low:high])
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
