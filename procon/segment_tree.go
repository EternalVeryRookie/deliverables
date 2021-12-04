package procon

import (
	"math"
	"sort"
)

type monoidUint64 struct {
	operate         func(x, y uint64) uint64
	identityElement uint64
}

func newMinOperateMonoid() monoidUint64 {
	return monoidUint64{
		operate: func(x, y uint64) uint64 {
			if x < y {
				return x
			}

			return y
		},
		identityElement: math.MaxUint64,
	}
}

func newSumOperateMonoid() monoidUint64 {
	return monoidUint64{
		operate: func(x, y uint64) uint64 {
			return x + y
		},
		identityElement: 0,
	}
}

//n以上の中で最小の2^kの値を出力
func calcMin2PowersMoreThan(n uint64) uint64 {
	i := 1
	for {
		powerOf2 := uint64(math.Pow(2, float64(i)))
		if n > powerOf2 {
			i++
		} else {
			return powerOf2
		}
	}
}

//len(arr)以上の中で最小の2^kのlengthを持つスライスを生成し、[0:len(arr)]にはarrの要素をコピーする。空いた要素にはpaddingをセットする。
func convertLengthMin2PowersMoreThan(arr []uint64, padding uint64) []uint64 {
	length := calcMin2PowersMoreThan(uint64(len(arr)))
	retArr := make([]uint64, length)
	for i := range retArr {
		if i < len(arr) {
			retArr[i] = arr[i]
		} else {
			retArr[i] = padding
		}
	}

	return retArr
}

func duplicateRange(x, y Range) *Range {
	if x.start <= y.start {
		if x.start+x.length <= y.start {
			return nil
		}

		length := (x.start + x.length) - y.start
		if y.length < length {
			length = y.length
		}
		return &Range{
			start:  y.start,
			length: length,
		}
	} else {
		return duplicateRange(y, x)
	}
}

type Range struct {
	start  int
	length int
}

func (r Range) isInclude(src Range) bool {
	rEnd := r.start + r.length
	srcEnd := src.start + src.length
	return src.start >= r.start && srcEnd <= rEnd
}

func (r Range) isEqual(src Range) bool {
	return r.start == src.start && r.length == src.length
}

type ISegmentNode interface {
	computeSolution(func(x, y uint64) uint64)
	getSolution() uint64
	getRange() Range
}

type SegmentTreeLeaf struct {
	value  uint64
	index  int
	parent *SegmentTreeNode
}

func (l *SegmentTreeLeaf) setValue(v uint64, opt func(x, y uint64) uint64) {
	l.value = v
	l.computeSolution(opt)
}

func (l *SegmentTreeLeaf) getSolution() uint64 {
	return l.value
}

func (l *SegmentTreeLeaf) computeSolution(opt func(x, y uint64) uint64) {
	l.parent.computeSolution(opt)
}

func (l *SegmentTreeLeaf) getRange() Range {
	return Range{start: l.index, length: 1}
}

type SegmentTreeNode struct {
	r          Range
	solution   uint64
	parent     *SegmentTreeNode
	leftChild  ISegmentNode
	rightChild ISegmentNode
}

func (n *SegmentTreeNode) computeSolution(opt func(x, y uint64) uint64) {
	n.solution = opt(n.leftChild.getSolution(), n.rightChild.getSolution())
	if n.parent != nil {
		n.parent.computeSolution(opt)
	}
}

func (n *SegmentTreeNode) getSolution() uint64 {
	return n.solution
}

func (n *SegmentTreeNode) getRange() Range {
	return n.r
}

func buildTree(parent *SegmentTreeNode, arr []uint64, start, length int) (root ISegmentNode, leafs []*SegmentTreeLeaf) {
	if length == 1 {
		self := &SegmentTreeLeaf{
			index:  start,
			value:  arr[start],
			parent: parent,
		}

		return self, []*SegmentTreeLeaf{self}
	}

	center := start + length/2
	self := &SegmentTreeNode{
		r:        Range{start: start, length: length},
		solution: 0,
		parent:   parent,
	}

	leftChild, leftLeafs := buildTree(self, arr, start, center-start)
	rightChild, rightLeafs := buildTree(self, arr, center, length-(center-start))

	self.leftChild = leftChild
	self.rightChild = rightChild

	return self, append(leftLeafs, rightLeafs...)
}

type SegmentTree struct {
	monoid monoidUint64
	root   ISegmentNode
	leafs  []*SegmentTreeLeaf
}

func NewSegmentTreeUint64(monoid monoidUint64, src []uint64) *SegmentTree {
	arr := convertLengthMin2PowersMoreThan(src, monoid.identityElement)
	root, leafs := buildTree(nil, arr, 0, len(arr))
	sort.Slice(leafs, func(i, j int) bool { return leafs[i].index < leafs[j].index })
	for _, l := range leafs {
		l.computeSolution(monoid.operate)
	}

	return &SegmentTree{
		monoid: monoid,
		root:   root,
		leafs:  leafs,
	}
}

func (t *SegmentTree) Set(value uint64, index int) {
	t.leafs[index].setValue(value, t.monoid.operate)
}

func (t *SegmentTree) Get(index int) uint64 {
	return t.leafs[index].value
}

func (t *SegmentTree) Query(query Range) uint64 {
	type subquery struct {
		n ISegmentNode
		q Range
	}
	sol := t.monoid.identityElement
	stack := []subquery{
		{
			n: t.root,
			q: query,
		},
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r := top.n.getRange()
		subQ := top.q
		if r.length == 1 || r.isEqual(subQ) {
			sol = t.monoid.operate(sol, top.n.getSolution())
		} else {
			node := top.n.(*SegmentTreeNode)
			leftChild := node.leftChild
			rightChild := node.rightChild
			if leftChild.getRange().isInclude(subQ) {
				stack = append(stack, subquery{
					n: leftChild,
					q: subQ,
				})
			} else if rightChild.getRange().isInclude(subQ) {
				stack = append(stack, subquery{
					n: rightChild,
					q: subQ,
				})
			} else {
				leftRange := duplicateRange(leftChild.getRange(), subQ)
				rightRange := duplicateRange(rightChild.getRange(), subQ)

				if leftRange != nil {
					stack = append(stack, subquery{n: leftChild, q: *leftRange})
				}

				if rightRange != nil {
					stack = append(stack, subquery{n: rightChild, q: *rightRange})
				}
			}
		}
	}

	return sol
}
