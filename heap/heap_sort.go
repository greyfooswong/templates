package heap

/**
在堆中，绝大多数操作都是通过根节点进行的

插入一个数{
size++
h[size] = num
up(size)
}
即将节点插入到size位置，然后使其上升

删除根节点{
h[1] = h[size]
size--
down(1)
}
将根节点的值覆盖为最后一个节点，然后缩减该树的size
之后对更新好的根节点进行down，维护其成为一个堆

删除某一元素{
h[k] = h[size]
size--
down(k) || up(k)
}
将被删除节点覆盖成最后一个节点，然后根据其比原来值的大小，进行down或up的维护

修改某一元素{
h[k] = x
down(k) || up(k)
直接将其值修改，之后根据其值和原来值的大小，进行down或up维护
}
*/

/**
根据题目给定的二叉树节点范围，决定这N
h数组是用来存储二叉树的，若当前节点是x，则其左儿子为2 * x，右儿子为2 * x + 1
size表示当前树的大小
*/
const N = 100010

var h [N]int
var size int

/**
首先将给定的所有数据存进树中
然后通过down操作，将当前树转化为堆，注意，只有堆顶元素是最小的，其他元素是无序的
*/
func make_heap(nums []int) {
	for i := range nums {
		h[i+1] = nums[i]
	}
	size = len(nums)
	for i := size / 2; i > 0; i-- {
		down(i)
	}
}

/**
down操作是将当前节点下沉
在这样处理之后，当前节点一定是和其左子节点和右子节点中最小的那个数
这样，就会将当前节点和最小节点的值交换
交换后，之前最小节点的值是最开始处理的节点，其仍旧可能不是最小的，因此要将其再次做down处理，保证在后续的查询和down操作中，严格保持着向下非递减的顺序
*/
/**
down操作中，首先记录当前节点的位置，然后将其与左右子节点做比较，记录这三个节点之间的最小值的编号，然后进行比对，如果最小值编号和当前节点编号不同
则交换当前节点和最小值节点的值，之后再对被交换的节点进行down处理
*/
func down(u int) {
	t := u
	if u*2 <= size && h[u*2] < h[t] {
		t = u * 2
	}
	if u*2+1 <= size && h[u*2] < h[t] {
		t = u*2 + 1
	}
	if u != t {
		h[u], h[t] = h[t], h[u]
		down(t)
	}
}

/**
findK函数是用来查找树中的第k小的节点，注意，查找过程中是不断删除根节点，并不断进行堆维护，以保证根节点是当前最小节点的，因此该操作会删除树中节点，且不可逆
如果想要使该操作可逆，可以记录下所有被删除的节点，之后将他们插入到该树中，并进行维护即可。
*/
func findK(k int) int {
	k--
	for k > 0 {
		k--
		h[1] = h[size]
		size--
		down(1)
	}
	return h[1]
}
