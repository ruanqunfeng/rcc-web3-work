package gobasework

// 0, 1, 2, 3, 4, 5, 6, 7, 8

// 反转字符串
func ReverseString(s []byte) {
	l := len(s) - 1
	for i := 0; i < l; i, l = i+1, l-1 {
		s[i], s[l] = s[l], s[i]
	}
}
