package arrays

func Sum(numbers []int) int {

	var result int
	for _, number := range numbers {
		result += number
	}
	return result

}

func SumAll(silcesOfnumbers ...[]int) (totalSums []int) {

	for _, slice := range silcesOfnumbers {
		totalSums = append(totalSums, Sum(slice))
	}
	return

}

func SumAllTails(silcesOfnumbers ...[]int) (totalTailSums []int) {

	for _, slice := range silcesOfnumbers {
		if len(slice) == 0 {
			totalTailSums = append(totalTailSums, 0)
		} else {
			tail := slice[1:]
			totalTailSums = append(totalTailSums, Sum(tail))
		}
	}
	return

}
