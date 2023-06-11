package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	gen1 := initGen(m, n)
	indices := gen1.gen()
	for _, pair := range indices {
		var v int = -1
		if head != nil {
			v = head.Val
			head = head.Next
		}
		res[pair.row][pair.col] = v
	}

	return res
}

type Gen struct {
	left   int
	right  int
	top    int
	bottom int
}

func initGen(m, n int) Gen {
	return Gen{
		bottom: m,
		right:  n,
	}
}

type Pair struct {
	row int
	col int
}

func (g *Gen) gen() []Pair {
	indices := make([]Pair, 0, g.bottom*g.right)
	dir := right

	for {
		if g.top == g.bottom || g.left == g.right {
			break
		}

		switch dir {
		case right:
			for i := g.left; i < g.right; i++ {
				indices = append(indices, Pair{
					row: g.top,
					col: i,
				})
			}
			g.top++
			dir = down
			// fmt.Println("after right", indices)
			continue
		case down:
			for i := g.top; i < g.bottom; i++ {
				indices = append(indices, Pair{
					row: i,
					col: g.right - 1,
				})
			}
			g.right--
			dir = left
			// fmt.Println("after down ", indices)
			continue
		case left:
			for i := g.right - 1; i >= g.left; i-- {
				indices = append(indices, Pair{
					row: g.bottom - 1,
					col: i,
				})
			}
			g.bottom--
			dir = up
			// fmt.Println("after left ", indices)
			continue
		case up:
			for i := g.bottom - 1; i >= g.top; i-- {
				indices = append(indices, Pair{
					row: i,
					col: g.left,
				})
			}
			g.left++
			dir = right
			// fmt.Println("after up   ", indices)
			continue
		}
	}

	return indices
}

const (
	right = iota
	down
	left
	up
)

func main() {
	spiralMatrix(4, 4, nil)
}
