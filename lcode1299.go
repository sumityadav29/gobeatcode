package main

func replaceElements(arr []int) []int {
	maxSoFar := -1

	for i := len(arr) - 1; i >= 0; i-- {
		currentElement := arr[i]
		arr[i] = maxSoFar

		maxSoFar = max(maxSoFar, currentElement)
	}

	return arr
}

// array elements can be mutated
// max function
