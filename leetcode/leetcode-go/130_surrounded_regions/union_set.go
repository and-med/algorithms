package main

type Point struct {
	X int
	Y int
}

type UnionSet struct {
	parent map[Point]Point
}

func (u *UnionSet) union(a Point, b Point) {
	aParent := u.getParent(a)
	bParent := u.getParent(b)

	if aParent != bParent {
		u.parent[b] = aParent
	}
}

func (u *UnionSet) getParent(p Point) Point {
	parent, ok := u.parent[p]

	if !ok {
		parent = p
	} else if parent != p {
		parent = u.getParent(parent)
		u.parent[p] = parent
	}

	return parent
}

const cross byte = 'X'
const oval byte = 'O'

func solve(board [][]byte) {
	u := UnionSet{make(map[Point]Point)}

	for i, row := range board {
		for j, val := range row {
			if val == oval && i+1 < len(board) && board[i+1][j] == oval {
				u.union(Point{i, j}, Point{i + 1, j})
			}
			if val == oval && j+1 < len(row) && board[i][j+1] == oval {
				u.union(Point{i, j}, Point{i, j + 1})
			}
		}
	}

	touchingEdge := make(map[Point]bool)
	n := len(board) - 1
	for i := 0; i < n; i++ {
		touchingEdge[u.getParent(Point{0, i})] = true
		touchingEdge[u.getParent(Point{i, 0})] = true
		touchingEdge[u.getParent(Point{n, i})] = true
		touchingEdge[u.getParent(Point{i, n})] = true
	}

	for i, row := range board {
		for j, val := range row {
			if val == oval && !touchingEdge[u.getParent(Point{i, j})] {
				board[i][j] = cross
			}
		}
	}
}

func main() {}
