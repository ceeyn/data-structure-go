package main

// [[5,10],[6,8],[1,5],[2,3],[1,10]]
// [1,5] [1,10] [2,3] [6,8] [5,10]
func subarraysWithKDistinct(nums []int, k int) int {
	return qihao_k(nums, k) - qihao_k(nums, k-1)
}

func qihao_k(nums []int, k int) int {
	left := 0
	ans := 0
	m := make(map[int]int)
	for k, v := range nums {
		m[v]++
		if len(m) > k {
			m[left]--
			left++
		}
		ans += k - left + 1
	}
	return ans
}
