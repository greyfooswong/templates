package two_part_graph

/**
无向图
将图染色，判定当前点是否被染色，如果未被染色，则将当前点染色为1，然后深度优先搜索，将其相连的点分别染色为2，继续向下搜索，直到遍历结束
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010
const M = 2 * N

var e, ne, h [M]int
var n, m, idx int
var color [N]int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func iniT() {
	for i := range h {
		h[i] = -1
	}
}

/**
将当前点染色
遍历和当前点相连的所有节点
如果当前节点没有染色，且之后染色都正确，则不处理
如果染色失败，则说明不是二分图
如果当前遍历节点染色并且与根节点颜色相同，则说明不是二分图
*/
func dfs(u, c int) bool {
	color[u] = c
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if color[j] == 0 {
			if !dfs(j, 3-c) {
				return false
			}
		} else if color[j] == c {
			return false
		}
	}
	return true
}

func main() {
	iniT()
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 20000*1024)
	sc.Buffer(buf, len(buf))
	sc.Scan()
	lines := strings.Split(sc.Text(), " ")
	n, _ = strconv.Atoi(lines[0])
	m, _ = strconv.Atoi(lines[1])
	for i := 0; i < m; i++ {
		sc.Scan()
		nums := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		add(a, b)
		add(b, a)
	}
	flag := false
	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			if !dfs(i, 1) {
				flag = true
				break
			}
		}
	}
	if flag {
		fmt.Printf("No")
	} else {
		fmt.Printf("Yes")
	}
}
