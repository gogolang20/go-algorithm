package graph

type Graph struct {
	nodes map[*Node]struct{}
	edges map[*Edge]struct{}
}

type Edge struct {
	weight int
	from   *Node
	to     *Node
}

type Node struct {
	value int

	nodes []*Node
	edges []*Edge

	in  int
	out int
}
