package xjoin

func Join(args ...[]interface{}) [][]interface{} {
	// Initialize a result array
	var result [][]interface{}

	// Return nothing if no slices were provided
	if len(args) == 0 {
		return [][]interface{}{}
		// Return a result with a single slice if only one was provided
	} else if len(args) == 1 {
		return append(result, args[0])
	}

	// instantiate the pools
	pools := [][]interface{}{}
	pools = append(pools, args...)

	// Make len(result) return 1 so that we can range over it
	result = append(result, []interface{}{})

	// Iterate over each of the slices
	// to make the cartesian product
	for _, pool := range pools {
		_result := result
		result = [][]interface{}{}
		for _, x := range _result {
			for _, y := range pool {
				result = append(result, append(x, y))
			}
		}
	}

	// Return the result
	return result
}
