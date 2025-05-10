package search

// BinarySearch 实现了一个二分查找算法。
// 返回一个查找的target在input的位置，不存在返回-1.
// 需要注意的是，二分查找要求输入是有序的。
func BinarySearch(input []int, target int) int {

	left := 0
	right := len(input) - 1

	var mid int

	for left <= right {

		mid = left + (right-left+1)/2

		switch {
		case input[mid] == target:
			return mid
		case input[mid] > target:
			right = mid - 1
		default:
			left = mid + 1
		}
	}

	return -1
}
