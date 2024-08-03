package sort

func quick_sort(nums []int, l, r int) {
	if l < r {
		pos_index := partition(nums, l, r)
		quick_sort(nums, l, pos_index-1)
		quick_sort(nums, pos_index+1, r)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[l]
	pre_l := l
	l++
	for l <= r {
		for l <= r && nums[l] <= pivot {
			l++
		}
		for l <= r && nums[r] >= pivot {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[pre_l], nums[r] = nums[r], nums[pre_l]
	return r
}
