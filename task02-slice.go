package homework

func reverse(input []int64) (result []int64) {
	res := make([]int64, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		res = append(res, input[i])
	}
	return res
}
