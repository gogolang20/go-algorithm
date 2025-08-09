package dp

func Sum(arr []int) []int {
	result := make([]int, len(arr))
	result[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		result[i] = result[i-1] + arr[i]
	}

	return result
}

func GetRangeSum(arr []int, l, r int) int {
	if l == 0 {
		return arr[r]
	}

	return arr[r] - arr[l-1]
}

/*
合并石子问题
*/

func mergeStones(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return mergeStonesProcess(Sum(arr), 0)
}

func mergeStonesProcess(arr []int, index int) int {

	return 0
}
