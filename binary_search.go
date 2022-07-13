package creek

// The BinarySearch function finds the index of a target value within a sorted stream.
func (s Stream[T]) BinarySearch(item T) int {
	return binarySearch(s.Array, 0, len(s.Array)-1, item)
}

func binarySearch[T Streamable](arr []T, left int, right int, n T) int {
	if left > right {
		return -1
	}

	midIndex := left + (right-left)/2

	if arr[midIndex] == n {
		return midIndex
	}

	if arr[midIndex] > n {
		return binarySearch(arr, left, midIndex-1, n)
	}

	return binarySearch(arr, midIndex+1, right, n)
}
