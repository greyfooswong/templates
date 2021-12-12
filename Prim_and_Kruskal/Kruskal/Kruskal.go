package Kruskal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
用于稀疏图
将边长从小到大排序
轮询当前最小的边，判断两个点是否处于一个点集之中，如果不在，则将该点放入同一个点集中（并查集处理），并将该边的代价加到res
当所有边遍历完后，如果仍然有点不在同一点集中，则说明没有最小生成树
否则，则最小生成树代价就是res
*/

const N = 200010

type edge struct {
	A, B, W int
}

//edges为切片
var edges []edge
var n, m int
var p [N]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find(a int) int {
	if p[a] != a {
		p[a] = find(p[a])
	}
	return p[a]
}

func main() {
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
		c, _ := strconv.Atoi(nums[2])
		edges = append(edges, edge{a, b, c})
	}
	//处理edges
	edges = edges[:m]
	//将edges从小到大排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})
	for i := range p {
		p[i] = i
	}
	res := 0
	cnt := 0
	for i := 0; i < m; i++ {
		a, b, c := edges[i].A, edges[i].B, edges[i].W
		a = find(a)
		b = find(b)
		//判断两个点是否在同一个点集之中，如果不在，则处理
		if a != b {
			p[a] = b
			res += c
			cnt++
		}
	}
	if cnt < n-1 {
		fmt.Printf("impossible")
	} else {
		fmt.Printf("%d\n", res)
	}

}
