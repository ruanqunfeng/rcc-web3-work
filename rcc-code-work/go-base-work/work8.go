package gobasework

import "slices"

// Merge 合并区间
// 给定一个区间的集合，请合并所有重叠的区间。
// 示例 1:
// 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出: [[1,6],[8,10],[15,18]]
// 解释: 区间 [1,3] 和 [2,6] 重叠, 所以合并为 [1,6].

func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	var res [][]int
	res = append(res, intervals[0])
	x := 0
	for _, arrs := range intervals {
		if res[x][1] >= arrs[0] {
			res[x][1] = max(arrs[1], res[x][1])
		} else {
			x++
			res = append(res, arrs)
		}
	}

	return res
}
