package goldtk

func dedupe[T comparable](first, second []T) []T {
	duplicates := make(map[T]struct{})
	merge := append(first, second...)

	var result []T

	for _, elem := range merge {
		if _, ok := duplicates[elem]; !ok {
			duplicates[elem] = struct{}{}
			result = append(result, elem)
		}
	}

	return result
}
