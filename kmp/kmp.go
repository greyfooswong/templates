package kmp

/**
传入两个字符串，一个模板串p，一个待匹配串s
返回值是每次s成功匹配p之后，【s中的（模板串的第一个位置）】的下标切片
*/

/**
在处理ne数组时，实际上是将模板串复制一份，两个都是模板串，然后处理出一个ne数组出来
*/
func kmp(pp, ss string) []int {
	res := []int{}
	p := conv(pp)
	s := conv(ss)
	ne := make([]int, len(pp)+10)
	for i, j := 2, 0; i <= len(pp); i++ {
		for j != 0 && p[i] != p[j+1] {
			j = ne[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		ne[i] = j
	}
	for i, j := 1, 0; i <= len(ss); i++ {
		for j != 0 && s[i] != p[j+1] {
			j = ne[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == len(pp) {
			res = append(res, i-len(pp))
			j = ne[j]
		}
	}
	return res
}

/**
处理字符串，将字符串前填充一个空格，减少边界判断
*/
func conv(word string) []byte {
	res := []byte{}
	res = append(res, ' ')
	for i := range word {
		res = append(res, word[i])
	}
	return res
}
