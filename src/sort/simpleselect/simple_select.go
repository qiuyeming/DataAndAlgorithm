package simpleselect

/*
	简单插入排序，省去了交换的步骤，性能高于冒泡
	最差是O(N^2)

*/

// Sort
func Sort(data []int) {
	for i := 0; i < len(data); i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
		}
	}
}
