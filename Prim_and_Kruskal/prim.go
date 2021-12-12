package Prim_and_Kruskal

/*
S:当前已经在联通块中的所有点的集合
1. dist[i] = inf
2. for n 次
    t<-S外离S最近的点
    利用t更新S外点到S的距离
    st[t] = true
n次迭代之后所有点都已加入到S中
联系：Dijkstra算法是更新到起始点的距离，Prim是更新到集合S的距离
*/

const N = 510
const INF = 0x3f3f3f3f

//邻接矩阵g存储所有边
//dist存储其他点到S的距离
var g [N][N]int
var st [N]bool
var n, m int
var dist [N]int

func iniT() {
	for i := range g {
		dist[i] = INF
		for j := range g[i] {
			g[i][j] = INF
		}
	}
}

//向图中加入权边，因为是无向图，所以要维护两次
func add_to_g(a, b, c int) {
	g[a][b] = min(g[a][b], c)
	g[b][a] = g[a][b]
}

//如果图不连通返回INF, 否则返回res
func Prim() int {
	res := 0
	for i := 0; i < n; i++ {
		t := -1

		//寻找离集合S最近的点
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		//判断是否连通，有无最小生成树
		if i != 0 && dist[t] == INF {
			return INF
		}
		if i != 0 {
			res += dist[t]
		}
		st[t] = true

		//更新最新S的权值和
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], g[t][j])
		}

	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
