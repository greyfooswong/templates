package trie

//根据题目给定的数量决定,字符串的最长长度
const N = 100010

var son [N][26]int
var cnt [N]int
var idx int

//向字典树中插入一个字符串
func insert(word string) {
	p := 0
	for i := range word {
		u := word[i] - 'a'
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

//查询字典树中该字符串的数量
func query(word string) int {
	p := 0
	for i := range word {
		u := word[i] - 'a'
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
}
