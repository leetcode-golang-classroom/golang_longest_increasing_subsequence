package sol

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLIS(nums []int) int {
	nLen := len(nums)
	dp := make([]int, nLen)
	res := 0
	for start := nLen - 1; start >= 0; start-- {
		dp[start] = 1
		for end := start + 1; end < nLen; end++ {
			// check possible max
			if nums[start] < nums[end] {
				dp[start] = max(dp[start], 1+dp[end])
			}
		}
		res = max(res, dp[start])
	}
	return res
}
