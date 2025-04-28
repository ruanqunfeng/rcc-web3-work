package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m
	for j := 0; j < n; {
		nums1[i] = nums2[j]
		i++
		j++
	}

	sort.Ints(nums1)
}