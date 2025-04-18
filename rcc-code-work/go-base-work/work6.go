package gobasework

// 给定一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
func mySqrt(x int) int {

	m := 0
	l, r := 0, x
	for l <= r {
		// 使用位运算来避免整数溢出
		// l + (r - l) / 2，这个作用是求中间值，因为l 和 r会变
		mid := l + (r-l)/2
		if mid*mid <= x {
			m = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return m
}
