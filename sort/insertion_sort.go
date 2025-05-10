package sort

// InsertionSort 插入排序。将输入input升序排序后返回。
// 内层循环每次把当前的值，和之前的有序序列依次比较，将当前值移动到合适的位置后，break内层循环。
func InsertionSort(input []int) []int {

	var tmp int

	for i := range len(input) {
		for j := i; j > 0; j-- {
			if input[j] < input[j-1] {
				tmp = input[j]
				input[j] = input[j-1]
				input[j-1] = tmp
			} else {
				break
			}
		}
	}
	return input
}
