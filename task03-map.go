package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	// make array to store keys
	keys := make([]int, 0, len(input))

	// write keys into array
	for k := range input {
		keys = append(keys, k)
	}
	// sort keys in array
	sort.Ints(keys)

	// make arrray to store values
	results := []string{}

	// store values in array
	for _, value := range keys {
		results = append(results, input[value])
	}
	return results
}
