package bs

/*
查找右边界
*/
func Find_right(q []int, k int) int {
	l, r := 0, len(q)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if q[mid] <= k {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if q[l] != k {
		return -1
	}
	return l
}
