package XYL

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
匈牙利算法
计算二分图的最大匹配
二分图两个点集为n1和n2
匹配指n1中的点只有一条边与n2中的点相连，且n2中也只有一条边与n1中的点相连
*/

const N = 510
const M = 100010

var e, ne, h, match [M]int

//n1为n1的点集数量，n2为n2的点集数量，m为两个点集之间的边数
var n1, n2, m, idx int
var st [N]bool

//每次遍历前将当前所有被匹配的n2点集中的点赋为false
func iniT1() {
	for i := range st {
		st[i] = false
	}
}

//初始化所有头节点
func iniT2() {
	for i := range h {
		h[i] = -1
	}
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

//判断当前n1点集的点是否有n2点集中的点可以匹配
func find(a int) bool {
	for i := h[a]; i != -1; i = ne[i] {
		j := e[i]
		//首先确定该n2点尚未被匹配
		if !st[j] {
			//将该n2点标记为被匹配
			st[j] = true
			//如果该n2点的值尚未被匹配，或者该节点已经被匹配了但是匹配其的n1点可以与其他点进行匹配，则匹配成功，并标记该n2点为该n1点的匹配点
			if match[j] == 0 || (find(match[j])) {
				match[j] = a
				return true
			}
		}
	}
	//匹配失败
	return false
}

func main() {
	iniT2()
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 20000*1024)
	sc.Buffer(buf, len(buf))
	sc.Scan()
	lines := strings.Split(sc.Text(), " ")
	n1, _ = strconv.Atoi(lines[0])
	n2, _ = strconv.Atoi(lines[1])
	m, _ = strconv.Atoi(lines[2])
	for i := 0; i < m; i++ {
		sc.Scan()
		nums := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		add(a, b)
	}
	res := 0
	//每次遍历之前将st全部赋为false，避免影响
	for i := 1; i <= n1; i++ {
		iniT1()
		if find(i) {
			res++
		}
	}
	fmt.Printf("%d", res)
}
