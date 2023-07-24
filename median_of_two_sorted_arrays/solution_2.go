package median_of_two_sorted_arrays

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		return float64(getKth(length/2, nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1)+getKth(length/2+1, nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1)) / 2
	} else {
		return float64(getKth(length/2+1, nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1))
	}
}

func getKth(k int, nums1 []int, l1, r1 int, nums2 []int, l2, r2 int) int {
	length1, length2 := r1-l1+1, r2-l2+1
	if length1 <= 0 {
		return nums2[l2+k-1]
	}
	if length2 <= 0 {
		return nums1[l1+k-1]
	}
	if k == 1 {
		return min(nums1[l1], nums2[l2])
	}

	m1 := min(r1, l1+k/2-1)
	m2 := min(r2, l2+k/2-1)

	if nums1[m1] < nums2[m2] {
		newL1 := m1 + 1
		newK := k - (newL1 - l1)
		return getKth(newK, nums1, newL1, r1, nums2, l2, r2)
	} else {
		newL2 := m2 + 1
		newK := k - (newL2 - l2)
		return getKth(newK, nums1, l1, r1, nums2, newL2, r2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
