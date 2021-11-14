package hash

/**
模拟哈希表
注意，在初始化时需要将h内的所有元素赋值为1
*/

/**
找到数据范围之后的第一个质数，该质数可以降低哈希冲突的概率
*/
const N = 100003

/**
h数组存储的是当前节点在哈希表中的位置，其值是idx
e存储的值
ne存储的是当前节点的下一个节点
idx指示的是当前e的存储位置，赋给h表示h[k]这一哈希槽指向的第一个节点是idx
*/
var h, e, ne [N]int
var idx int

/**
将当前值存储
令当前节点指向h[k]指向的节点
令h[k]节点指向当前节点
增加存储的数量idx++
*/
func insert(x int) {
	k := (x%N + N) % N
	e[idx] = x
	ne[idx] = h[k]
	h[k] = idx
	idx++
}

/**
首先求当前x的哈希值，然后去查找对应的节点，从对应的节点中开始判断
如果查找的当前节点不是要找的值，则继续向其后的节点查找
直到查找到或为空为止
*/
func find(x int) bool {
	k := (x%N + N) % N
	for i := h[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return true
		}
	}
	return false
}
