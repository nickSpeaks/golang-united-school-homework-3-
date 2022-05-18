package homework

func sortMapValues(input map[int]string) (result []string) {
		// make array to store keys
		keys := make([]int, 0, len(input))

		for k := range input {
			keys = append(keys, k)
		}
		// sort keys in array
		sort.Ints(keys)
	
		// make arrray to store values
		result := []string{}
	
		// store values in array
		for _, value := range keys {
			result = append(result, input[value])
	return
}
