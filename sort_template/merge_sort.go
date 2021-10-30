package sort

//temp根据数据范围来创建
var temp = make([]int, int(1e6)+10)

func Merge_sort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	Merge_sort(q, l, mid)
	Merge_sort(q, mid+1, r)
	k, i, j := 0, l, mid+1

	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			k++
			i++
		} else {
			temp[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}
	for i, j = l, 0; i <= r; {
		q[i] = temp[j]
		i++
		j++
	}
}
