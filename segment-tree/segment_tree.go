package segment_tree

type SegmentTree struct {
	data   []int // 原始数组
	tree   []int // 下标从 0 开始构建
	lazy   []int
	change []int
	update []bool
}

func NewSegmentTree(arr []int) *SegmentTree {
	length := len(arr)

	st := &SegmentTree{
		data:   make([]int, 0, length),
		tree:   make([]int, length<<2),
		lazy:   make([]int, length<<2),
		change: make([]int, length<<2),
		update: make([]bool, length<<2),
	}

	for index := range arr {
		st.data = append(st.data, arr[index])
	}

	st.Build(0, 0, length-1)

	return st
}

func (st *SegmentTree) Build(treeIndex, left, right int) {
	if left == right {
		st.tree[left] = st.data[left]
		return
	}

	mid := left + (right-left)>>1
	leftIndex := treeIndex*2 + 1
	rightIndex := treeIndex*2 + 2

	st.Build(leftIndex, left, mid)
	st.Build(rightIndex, mid+1, right)

	st.tree[treeIndex] = st.tree[leftIndex] + st.tree[rightIndex]
}

func (st *SegmentTree) pushDown(treeIndex, leftSize, rightSize int) {
	leftIndex := treeIndex*2 + 1
	rightIndex := treeIndex*2 + 2

	if st.update[treeIndex] {
		st.update[leftIndex] = true
		st.update[rightIndex] = true

		st.change[leftIndex] = st.lazy[leftIndex]
		st.change[rightIndex] = st.lazy[rightIndex]

		st.tree[leftIndex] += st.lazy[treeIndex] * leftSize
		st.tree[rightIndex] += st.lazy[treeIndex] * rightSize

		st.lazy[leftIndex] = 0
		st.lazy[rightIndex] = 0

		st.update[treeIndex] = false
	}
	if st.lazy[treeIndex] != 0 {
		st.lazy[leftIndex] = st.lazy[treeIndex]
		st.lazy[rightIndex] = st.lazy[treeIndex]

		st.tree[leftIndex] += st.lazy[treeIndex] * leftSize
		st.tree[rightIndex] += st.lazy[treeIndex] * rightSize

		st.lazy[treeIndex] = 0
	}
}

func (st *SegmentTree) Query(rangeLeft, rangeRight int) int {
	return st.QueryInTree(0, 0, len(st.data)-1, rangeLeft, rangeRight)
}

func (st *SegmentTree) QueryInTree(treeIndex, left, right, rangeLeft, rangeRight int) int {
	if rangeLeft <= left && right <= rangeRight {
		return st.tree[treeIndex]
	}

	mid := left + (right-left)>>1
	st.pushDown(treeIndex, mid-left+1, right-mid)

	leftIndex := treeIndex*2 + 1
	rightIndex := treeIndex*2 + 2

	result := 0
	if rangeLeft <= mid {
		result += st.QueryInTree(leftIndex, left, mid, rangeLeft, rangeRight)
	}
	if rangeRight > mid {
		result += st.QueryInTree(rightIndex, mid+1, right, rangeLeft, rangeRight)
	}

	return result
}

func (st *SegmentTree) Update(rangeLeft, rangeRight, val int) {
	st.UpdateInTree(0, 0, len(st.data)-1, rangeLeft, rangeRight, val)
}

func (st *SegmentTree) UpdateInTree(treeIndex, left, right, rangeLeft, rangeRight, val int) {
	if rangeLeft <= left && right <= rangeRight {
		st.change[treeIndex] = val
		st.update[treeIndex] = true
		st.tree[treeIndex] = val * (right - left + 1)
		st.lazy[treeIndex] = 0
		return
	}

	mid := left + (right-left)>>1
	st.pushDown(treeIndex, mid-left+1, right-mid)

	leftIndex := treeIndex*2 + 1
	rightIndex := treeIndex*2 + 2
	if rangeLeft <= mid {
		st.UpdateInTree(leftIndex, left, mid, rangeLeft, rangeRight, val)
	}
	if rangeRight > mid {
		st.UpdateInTree(rightIndex, mid+1, right, rangeLeft, rangeRight, val)
	}

	st.tree[treeIndex] = st.tree[leftIndex] + st.tree[rightIndex]
}

func (st *SegmentTree) Add(rangeLeft, rangeRight, val int) {
	st.AddInTree(0, 0, len(st.data)-1, rangeLeft, rangeRight, val)
}

func (st *SegmentTree) AddInTree(treeIndex, left, right, rangeLeft, rangeRight, val int) {
	if rangeLeft <= left && right <= rangeRight {
		st.tree[treeIndex] += val * (right - left + 1)
		st.lazy[treeIndex] += val
		return
	}

	mid := left + (right-left)>>1
	st.pushDown(treeIndex, mid-left+1, right-mid)

	leftIndex := treeIndex*2 + 1
	rightIndex := treeIndex*2 + 2
	if rangeLeft <= mid {
		st.AddInTree(leftIndex, left, mid, rangeLeft, rangeRight, val)
	}
	if rangeRight > mid {
		st.AddInTree(rightIndex, mid+1, right, rangeLeft, rangeRight, val)
	}

	st.tree[treeIndex] = st.tree[leftIndex] + st.tree[rightIndex]
}
