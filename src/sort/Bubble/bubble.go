package Bubble
/*
最好的情况是做n-1次比较，不移位 O(N)
最坏的情况是完全逆序，O(N^2)
*/


// BubbleSort
// 两两交换
func BubbleSort(data []int) {
	var flag = true
	for i := 0; i < len(data) && flag; i++ {
		flag = false
		for j := len(data)-1; j > i; j-- {
			if data[j] < data[j-1] {
				data[j-1], data[j] = data[j], data[j-1]
				flag = true
			}
		}
	}
}



