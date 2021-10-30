package prefix_sum

//为了防止溢出，使用int64返回
//注意，这里是考虑边界的情况，即S0 = 0， 而S1才是左边第一个元素的值
func Prefix_sum_one(q []int) []int64 {
	res := make([]int64, len(q)+1)
	for i := 1; i <= len(q); i++ {
		res[i] = res[i-1] + int64(q[i-1])
	}
	return res
}

//为了防止溢出，使用int64返回
//注意，这里是考虑边界的情况，即S0 = 0，而s1才是左上角第一个元素的值
func Prefix_sum_two(q [][]int) [][]int64 {
	res := make([][]int64, len(q)+1)
	for k := range res {
		res[k] = make([]int64, len(q[0])+1)
	}
	for i := 1; i <= len(q); i++ {
		for j := 1; j < len(q[0]); j++ {
			res[i][j] = res[i-1][j] + res[i][j-1] - res[i-1][j-1] + int64(q[i-1][j-1])
		}
	}
	return res
}
