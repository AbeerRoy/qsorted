package qsorted

func qsorted(arr []int) []int {
	var pivot = len(arr) - 1

	if pivot <= 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {

		if arr[i] > arr[pivot] {
			var temp = 0
			temp = arr[i]
			arr[i] = arr[pivot]
			arr[pivot] = temp

		}

	}

	pivot--
	qsorted(arr[:pivot+1])

	return arr
}
