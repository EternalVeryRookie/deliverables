package procon

type Point2D struct {
	x, y int64
}

type KdTree struct {
	points []Point2D
}

type KdTreeNode struct {
	min, max Point2D
}

func createKdTree(data [][2]int64) {

}
