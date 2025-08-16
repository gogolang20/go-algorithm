package unionfind

type Node struct {
}

type UnionFindMap struct {
	parent map[*Node]*Node
	size   map[*Node]int
	sets   int
}
