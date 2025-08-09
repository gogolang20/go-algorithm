package dp

func canPartition1(nums []int) bool {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	if sum&1 != 0 {
		return false
	}
	aim := sum >> 1

	return process(nums, 0, aim)
}

func process(nums []int, index, rest int) bool {
	if index == len(nums) {
		if rest == 0 {
			return true
		}
		return false
	}

	return process(nums, index+1, rest) || process(nums, index, rest-nums[index])
}

// canPartition 动态规划，空间压缩
func canPartition(nums []int) bool {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	if sum&1 != 0 {
		return false
	}
	aim := sum >> 1

	dp := make([]bool, aim+1)

	dp[0] = true
	for index := len(nums) - 1; index >= 0; index-- {
		for rest := aim; rest >= nums[index]; rest-- {
			dp[rest] = dp[rest] || dp[rest-nums[index]]
		}
	}

	return dp[aim]
}
