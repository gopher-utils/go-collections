package sort

import m "linq/types"

func partition[V m.Number](arr []V, low, high int) ([]V, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort[V m.Number](arr []V, low, high int) []V {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func Init[V m.Number](arr []V) []V {
	return quickSort(arr, 0, len(arr)-1)
}
