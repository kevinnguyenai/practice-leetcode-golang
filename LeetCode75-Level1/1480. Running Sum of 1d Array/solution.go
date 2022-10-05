package runningSum

func runningSum(nums []int) []int {
	result := []int{}
	result = append(result, nums[0])

	for i := 1; i < len(nums); i++ {
		result = append(result, result[i-1]+nums[i])
	}
	return result

}
