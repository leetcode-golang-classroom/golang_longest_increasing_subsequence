package sol

func lengthOfLISv1(nums []int) int {
	nLen := len(nums)
	seq := []int{nums[0]}
	var binarySearch = func(start, end, target int) int {
		left, right := start, end
		for left <= right {
			mid := (left + right) / 2
			if target == seq[mid] {
				return mid
			}
			if target > seq[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	for idx := 1; idx < nLen; idx++ {
		sLen := len(seq)
		insertIdx := binarySearch(0, sLen-1, nums[idx])
		if insertIdx == sLen {
			seq = append(seq, nums[idx])
		} else {
			seq[insertIdx] = nums[idx]
		}
	}
	return len(seq)
}
