package dijikstra

//如果图中有自环，只要是正的一律抛弃
//如果图中有重边，则选取最小的那个边记录下来

const N = 510
const INF = 0x3f3f3f3f

//g是保存a到b的距离的邻接矩阵
var g [N][N]int
var dist [N]int
var st [N]bool

var n, m int

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 初始化，将邻接矩阵和距离都赋为无穷大
func iNIT() {
	for i := range g {
		for j := range g[i] {
			g[i][j] = INF
		}
	}
	for i := range dist {
		dist[i] = INF
	}
}

func dijikstra() int {
	//从哪个点开始，就将哪个点赋为0
	dist[1] = 0
	for i := 0; i < n; i++ {
		t := -1
		//第一次遍历，将所有点中，除已经遍历过的点以外的，所有距离起点最小的点筛选出来
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}
		//将之前遍历找到的最小dist的点标记为已经遍历
		st[t] = true
		//第二次遍历，将所有点的dist与当前点到这个点的距离处理，确保每次遍历之前，所有的dist都已经是当前情况下的最小结果
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+g[t][j])
		}
	}
	//检查是否能到达点n，不能即为INF无穷大
	if dist[n] == INF {
		return -1
	}
	return dist[n]
}
