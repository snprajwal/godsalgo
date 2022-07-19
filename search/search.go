package search

// Perform linear search on arr for key
func Linear(arr []int64, key int64) int {
	for i, ele := range arr {
		if ele == key {
			return i
		}
	}
	return -1
}

// Perform iterative binary search on arr for key
func Binary(arr []int64, key int64) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		switch {
		case arr[mid] == key:
			return mid
		case arr[mid] > key:
			r = mid - 1
		case arr[mid] < key:
			l = mid + 1
		}
	}
	return -1
}

// Perform recursive binary search on arr for key
func BinaryRecursive(arr []int64, key int64) int {
	return binaryRecursive(arr, key, 0, len(arr)-1)
}

func binaryRecursive(arr []int64, key int64, l, r int) int {
	if l <= r {
		mid := l + (r-l)/2
		switch {
		case arr[mid] == key:
			return mid
		case arr[mid] > key:
			return binaryRecursive(arr, key, l, mid-1)
		case arr[mid] < key:
			return binaryRecursive(arr, key, mid+1, r)
		}
	}
	return -1
}
