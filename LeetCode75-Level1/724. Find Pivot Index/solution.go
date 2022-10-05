package pivotIndex

func pivotIndex(nums []int) int {
	var lsum, rsum int

	for i := 0; i < len(nums); i++ {
		rsum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		lsum += nums[i]
		if lsum == rsum {
			return i
		}
		rsum -= nums[i]
	}
	return -1
}
