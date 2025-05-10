package sort

// BubbleSort 冒泡排序。将输入input升序排序后返回。
func BubbleSort(input []int) []int {

	var tmp int

	for i := range len(input) {
		for j := range len(input) - i - 1 {

			if input[j] > input[j+1] {
				tmp = input[j]
				input[j] = input[j+1]
				input[j+1] = tmp
			}
		}
	}

	return input
}
