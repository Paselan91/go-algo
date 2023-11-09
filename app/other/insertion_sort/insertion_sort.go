package insertion_sort

func InsertionSort(arrays []int) []int {
	var j int
	var tmp int

	for i := range arrays {
		tmp = arrays[i]
		j = i - 1
		for j >= 0 && arrays[j] > tmp {
			arrays[j], arrays[j+1] = swap(arrays[j], arrays[j+1])
			j--
		}
		arrays[j+1] = tmp
	}

	return arrays
}

func swap(a int, b int) (int, int) {
	return b, a
}
