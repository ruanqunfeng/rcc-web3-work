package gobasework

/**
以输入[1,2,3]为例：

初始调用backtrack([1,2,3], 0, res)

start=0，循环i从0到2

i=0:

交换0和0（实际不变）

递归backtrack([1,2,3], 1, res)

start=1，循环i从1到2

i=1:

交换1和1（不变）

递归backtrack([1,2,3], 2, res)

start=2，循环i从2到2

i=2:

交换2和2（不变）

递归backtrack([1,2,3], 3, res)

保存排列[1,2,3]

i=2:

交换1和2位置：[1,3,2]

递归backtrack([1,3,2], 2, res)

保存排列[1,3,2]

i=1:

交换0和1位置：[2,1,3]

递归处理...

i=2:

交换0和2位置：[3,2,1]

递归处理...

最终生成所有6种排列。
**/
func permute(nums []int) [][]int {
	var res [][]int
	backtrack(nums, 0, &res)
	return res
}

func backtrack(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrack(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}

}
