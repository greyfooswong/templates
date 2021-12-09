package b_f

//可以存在负权边
//有边数限制的单源最短路，如果没有边数限制且存在负权边则边数限制为边的数量，即m
//可以判断是否存在负环，即根据单源最短路所有边走完之后，如果再次循环时，仍然有dist会更新，则说明存在负环

const N = 510
const M = 10010
const INF = 0x3f3f3f3f

//n个点，m条边，最多走k条边到
var n, m, k int

//dist用来存储当前的最小值，而backup用来备份上次的最小值，避免覆盖
var dist, backup [N]int

//边的结构体，A到B的权值为W的边
type edge struct {
	A, B, W int
}

//所有边的数组
var edges [M]edge

//将所有距离初始化为正无穷
func iNit() {
	for i := range dist {
		dist[i] = INF
	}
}

func b_f() int {
	//将1初始化为0，即从1开始
	dist[1] = 0
	for i := 0; i < k; i++ {
		//将上次的dist给备份下来
		for j := range dist {
			backup[j] = dist[j]
		}
		//遍历所有的边，将其对应的b做一次更新，这次更新的比较对象是目前的dist大小和上次遍历之后的backup到a的值
		for j := 0; j < m; j++ {
			a, b, w := edges[j].A, edges[j].B, edges[j].W
			dist[b] = min(dist[b], backup[a]+w)
		}
	}
	//如果遍历完后，dist【n】依旧是个正无穷的情况，则说明无法到达
	if dist[n] > INF/2 {
		return -INF
	}
	//否则返回到n的最小边
	return dist[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
