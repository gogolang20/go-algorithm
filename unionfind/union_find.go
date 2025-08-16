package unionfind

type UnionFind struct {
	parent []int
	help   []int
	size   []int
	sets   int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, size),
		help:   make([]int, size),
		size:   make([]int, size),
		sets:   size,
	}

	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	hi := 0
	for uf.parent[x] != x {
		uf.help[hi] = x
		hi++
		x = uf.parent[x]
	}

	for hi--; hi >= 0; hi-- {
		uf.parent[uf.help[hi]] = x
	}

	return x
}

func (uf *UnionFind) Union(x, y int) {
	xf := uf.Find(x)
	yf := uf.Find(y)
	if xf == yf {
		return
	}

	if uf.size[xf] > uf.size[yf] {
		uf.parent[yf] = xf
		uf.size[xf] += uf.size[yf]
	} else {
		uf.parent[xf] = yf
		uf.size[yf] += uf.size[xf]
	}
	uf.sets--
}

func (uf *UnionFind) IsSameSet(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
