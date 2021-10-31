package bitwise

//判断数字n倒数第k位是不是1

func Check(n, k int) bool {
	for k > 0 {
		k--
		n >>= 1
	}
	return (n & 1) == 1
}
