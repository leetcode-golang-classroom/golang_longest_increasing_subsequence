# golang_longest_increasing_subsequence

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A **subsequence** is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

## Examples

**Example 1:**

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

```

**Example 2:**

```
Input: nums = [0,1,0,3,2,3]
Output: 4

```

**Example 3:**

```
Input: nums = [7,7,7,7,7,7,7]
Output: 1

```

**Constraints:**

- `1 <= nums.length <= 2500`
- `104 <= nums[i] <= 104`

**Follow up:** Can you come up with an algorithm that runs in `O(n log(n))` time complexity?

## 解析

給定一個整數陣列 nums

要求寫一個演算法算出最長的嚴格遞增子序列長度

nums 的嚴格遞增子序列 sequences 需要符合幾個特性

1  sequences[i] < sequences[j] if i < j

2 sequences[i] 與 sequences[j] 在原本 nums 的順序也是跟 nums 一樣

舉一個例子來思考如何計算

假設 nums: [1,2,4,3]

可以發現以下決策樹

![](https://i.imgur.com/L2y8319.png)

假設以 dp[i] 代表從 i index 開始包含 nums[i] 最大子序列長度

則發現 對所有 dp[i] = max(1, 1+dp[i+1], 1+dp[i+2], … , 1+dp[len(nums)-1]) 

這樣就可以避免重複運算

從最後一個開始位置檢查

每次檢查都要從位置 i 檢查到 len(nums)-1 所以 O(n)

因為 從每個 index 開始都要檢查一次 所以 O(n)

所以時間複雜度會 O($n^2$)

要把時間複雜度優化到 O(nlogn)

現在以順向方式蒐集子序列

思考一下如何建構最小子序列

建構一個陣列用來保存目收集到的子序列保持嚴格遞增

每次考慮某一個 nums[i] 時先透過 binarySearch 檢查該 nums[i] 目前子序列的位置

假設這個位置是 len(nums) 代表 nums[i] 會是最大值

假設這個值 < len(nums) 則代表不影響長度 所以更新該位置為 nums[i]

![](https://i.imgur.com/GIJ4CJY.png)

## 程式碼
```go
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

```
## 困難點

1. 發現這個遞迴子問題關係不是很直觀
2. 要優化到 O(nlogn) 需要用到一些特性

## Solve Point

- [x]  建立一個長度為len(nums) 的陣列 dp 用來儲存每個開始位置的最大子序列長度
- [x]  然後逐步透過 dp[i] = max(1, 1+dp[i+1], 1+dp[i+2], … , 1+dp[len(nums)-1]) 累計出最大值