package gobasework

func RemoveDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	num := 1
	for x, y := 0, 1; y < l; y++ {
		if nums[x] != nums[y] {
			num++
			nums[x+1] = nums[y]
			x++
		}
	}

	return num
}
