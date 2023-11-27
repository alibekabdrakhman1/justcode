package two_sum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	ans := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		val, ok := m[target-nums[i]]
		if ok {
			ans[0] = i
			ans[1] = val
		} else {
			m[nums[i]] = i
		}
	}
	return ans
}
