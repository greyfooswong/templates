package floyd

//多源最短路，即可以求多个节点到其他某个节点的最短路径
//不能存在负权边
//在初始化时，只保留最短的边的距离

const N = 210
const INF = 0x3f3f3f3f

var n, m, k int
var g [N][N]int

//三重循环，获取从每个点到其他点的最短路径
func floyd() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for s := 1; s <= n; s++ {
				g[j][s] = min(g[j][s], g[j][i]+g[i][s])
			}
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//初始化邻接矩阵，将所有自环设为0，其他边权值设为无穷大
func ini() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				g[i][j] = 0
			} else {
				g[i][j] = INF
			}
		}
	}
}
