package dfs_and_bfs

const (
	N = 100010
	M = 2 * N
)

var e [N]int
var h [N]int
var ne [M]int
var st [N]bool
var idx int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs(u int) {
	st[u] = true
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if !st[j] {
			dfs(j)
		}
	}
}
