package sort

// SelectionSort 选择排序。将输入input升序排序后返回。
// 每趟循环找当前最小的值，然后与当前起始查找的位置交换。
func SelectionSort(input []int) []int {

	var tmp, minIndex int
	for i := range len(input) {
		minIndex = i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}

		tmp = input[i]
		input[i] = input[minIndex]
		input[minIndex] = tmp
	}

	return input
}
