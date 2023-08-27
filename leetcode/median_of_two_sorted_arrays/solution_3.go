package median_of_two_sorted_arrays

import "math"

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		return float64(getK(nums1, nums2, l/2)+getK(nums1, nums2, l/2+1)) / 2.0
	} else {
		return float64(getK(nums1, nums2, l/2+1))
	}
}

func getK(nums1 []int, nums2 []int, k int) int {
	if k == 1 {
		if len(nums1) > 0 && len(nums2) > 0 {
			min, max := nums1[0], nums2[0]
			if min > max {
				min, max = max, min
			}
			return min
		} else if len(nums1) > 0 {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	mid := k / 2
	p1, p2 := len(nums1)-1, len(nums2)-1
	if mid-1 < len(nums1) {
		p1 = mid - 1
	}
	if mid-1 < len(nums2) {
		p2 = mid - 1
	}

	var v1, v2 int
	if p1 < 0 {
		v1 = math.MaxInt
	} else {
		v1 = nums1[p1]
	}
	if p2 < 0 {
		v2 = math.MaxInt
	} else {
		v2 = nums2[p2]
	}

	if v1 <= v2 {
		nums1 = nums1[p1+1:]
		k -= p1 + 1
	} else {
		nums2 = nums2[p2+1:]
		k -= p2 + 1
	}

	return getK(nums1, nums2, k)
}
