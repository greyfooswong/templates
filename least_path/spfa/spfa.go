package spfa

/**
将起点距离更新为0，并将其入队，标记其为已加入的状态
之后每次从队头取出一个点，判断当前点所有出边是否需要修改，如果需要，则将该点加入到队列中，并标记其为已加入队列的状态
*/
const N = 100010
const INF = 0x3f3f3f3f

var e, ne, w, h, dist, cnt [N]int
var st [N]bool
var n, m, idx int

//初始化时，将表头赋为-1，dist赋为无穷大
func ini() {
	for i := range h {
		h[i] = -1
		dist[i] = INF
	}
}

//采用邻接表存图，c为边权
func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func spfa() int {
	//从点1到点n，因此dist[1] = 0
	dist[1] = 0
	q := []int{}
	//将起点入队
	q = append(q, 1)
	//标记起点已经入队
	st[1] = true
	for len(q) > 0 {
		//将队头出队，并标记一下
		t := q[0]
		q = q[1:]
		st[t] = false
		//从t开始，深度遍历所有节点，确保其目前处于一个最小的状态，如果不是，则将其入队，同时修改其dist
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
				if !st[j] {
					st[j] = true
					q = append(q, j)
				}
			}
		}
	}
	if dist[n] == INF {
		return INF
	}
	return dist[n]
}

//spfa判断是否存在负环
func spfa_check_has_circle() bool {

	//判断负环无法确定负环从那个点开始，因此要将所有点全部入队
	q := []int{}
	for i := 1; i <= n; i++ {
		q = append(q, i)
		st[i] = true
	}

	for len(q) > 0 {
		//将队头出队，并标记一下
		t := q[0]
		q = q[1:]
		st[t] = false
		//从t开始，深度遍历所有节点，确保其目前处于一个最小的状态，如果不是，则将其入队，同时修改其dist
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]

				//判断负环增加代码如下
				cnt[j] = cnt[t] + 1
				if cnt[j] >= n {
					return true
				}

				if !st[j] {
					st[j] = true
					q = append(q, j)
				}
			}
		}
	}
	return false
}
