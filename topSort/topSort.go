package topSort

//根据题目给定的范围而定
const N = 100010

var e, ne, h, q, d [N]int

//题目给定的n是图中节点数量，m是边的个数
var n, m, idx int

//经典存图
//h是标识当前头节点的位置
//e是标识当前要插入的节点
//ne是当前节点的下一个节点
func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

//拓扑排序
func topSort() bool {
	hh, tt := 0, -1
	//用数组模拟队列，首先将所有入度为0的节点加入队列中，数组d为标识节点被指向的边的数量
	for i := 1; i <= n; i++ {
		if d[i] == 0 {
			tt++
			q[tt] = i
		}
	}
	//将节点从图中删除后，所有与该节点相连的节点的入度减1，如果出现在减1后入度为0的节点，则将该节点加入队列中
	for hh <= tt {
		t := q[hh]
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			d[j]--
			if d[j] == 0 {
				tt++
				q[tt] = j
			}
		}
	}
	//最终只需要判断是否所有节点均入队即可，同时，如果所有节点均入队了，则当前队列中的数据顺序即为一个拓扑序
	return tt == n-1
}
