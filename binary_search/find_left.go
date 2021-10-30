package bs

// 查找左边界
func Find_left(q []int, k int) int {
	l, r := 0, len(q)-1
	for l < r {
		mid := (l + r) >> 1
		if q[mid] >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if q[l] != k {
		return -1
	}
	return l
}
