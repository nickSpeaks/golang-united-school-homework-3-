package homework

func average(input [15]float32) (result float32) {
	result = 0
	for _, n := range input {
		result += n
	}
	result /= 15
	return
}
