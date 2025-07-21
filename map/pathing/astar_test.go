package pathing

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type aStarTestNode struct {
	walkable bool
	index    int
	cost     int
}

func (n *aStarTestNode) GetPathCost() int {
	return n.cost
}

func (n *aStarTestNode) GetIndex() int {
	return n.index
}

func (n *aStarTestNode) IsWalkable() bool {
	return n.walkable
}

type aStarTestArena struct {
	m      []aStarTestNode
	height int
	width  int
}

func (m *aStarTestArena) GetApproximateDistance(from, to int) float64 {
	fromX, fromY := m.coordsFromIndex(from)
	toX, toY := m.coordsFromIndex(to)
	return math.Abs(float64(toX-fromX)) + math.Abs(float64(toY-fromY))
	// return math.Sqrt(math.Pow(math.Abs(float64(toX-fromX)), 2) + math.Pow(math.Abs(float64(toY-fromY)), 2))
}

func (m *aStarTestArena) GetNodeAtIndex(idx int) INode {
	return &m.m[idx]
}

func (m *aStarTestArena) GetNeighbors(idx int) []INode {
	x, y := m.coordsFromIndex(idx)

	coords := []struct {
		x int
		y int
	}{
		{
			x: x + 1,
			y: y,
		},
		{
			x: x - 1,
			y: y,
		},
		{
			x: x,
			y: y + 1,
		},
		{
			x: x,
			y: y - 1,
		},
	}
	res := []INode{}
	for _, node := range coords {
		if m.isWithinMap(node.x, node.y) {
			res = append(res, &m.m[m.indexFromCoords(node.x, node.y)])
		}
	}
	return res
}

func (m *aStarTestArena) indexFromCoords(x, y int) int {
	return y*m.width + x
}

func (m *aStarTestArena) coordsFromIndex(idx int) (int, int) {
	y := idx / m.width
	x := idx % m.width
	return x, y
}

func (m *aStarTestArena) isWithinMap(x int, y int) bool {
	if x < 0 || x > m.width-1 || y < 0 || y > m.height-1 {
		return false
	} else {

		return true
	}
}

func TestGetNeighbors(t *testing.T) {
	arena := aStarTestArena{
		m:      make([]aStarTestNode, 0),
		height: 5,
		width:  5,
	}

	data := []int{
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
	}
	for i := range data {
		arena.m = append(arena.m, aStarTestNode{
			walkable: data[i] == 1,
			index:    i,
			cost:     1,
		})
	}

	// indexArena := [][]int{
	// 	{0, 1, 2, 3, 4},
	// 	{5, 6, 7, 8, 9},
	// 	{10, 11, 12, 13, 14},
	// 	{15, 16, 17, 18, 19},
	// 	{20, 21, 22, 23, 24},
	// }
	tests := []struct {
		index             int
		expectedNeighbors []int
	}{
		{
			index:             6,
			expectedNeighbors: []int{1, 5, 7, 11},
		},
		{
			index:             19,
			expectedNeighbors: []int{14, 18, 24},
		},
		{
			index:             21,
			expectedNeighbors: []int{20, 22, 16},
		},
		{
			index:             5,
			expectedNeighbors: []int{0, 6, 10},
		},
		{
			index:             2,
			expectedNeighbors: []int{1, 3, 7},
		},
	}
	for i := range tests {
		test := tests[i]
		neighbors := arena.GetNeighbors(test.index)
		checkMap := make(map[int]struct{})
		for j := range neighbors {
			checkMap[neighbors[j].GetIndex()] = struct{}{}
		}
		require.Len(t, checkMap, len(test.expectedNeighbors), fmt.Sprintf("index: %d", test.index))
		for _, expected := range test.expectedNeighbors {
			delete(checkMap, expected)
		}
		require.Empty(t, checkMap)
	}

}

func TestCoordsFromIndex(t *testing.T) {
	arena := aStarTestArena{
		m:      make([]aStarTestNode, 0),
		height: 5,
		width:  5,
	}

	indexArena := []struct {
		x int
		y int
	}{
		{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
		{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1},
		{0, 2}, {1, 2}, {2, 2}, {2, 3}, {2, 4},
		{0, 3}, {1, 3}, {2, 3}, {3, 3}, {3, 4},
		{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
	}
	tests := []struct {
		index int
	}{
		{
			index: 6,
		},
		{
			index: 24,
		},
		{
			index: 20,
		},
	}
	for i := range tests {
		test := tests[i]
		x, y := arena.coordsFromIndex(test.index)
		require.Equal(t, indexArena[test.index].x, x, "comparing x")
		require.Equal(t, indexArena[test.index].y, y, "comparing y")
	}
}

func TestIndexFromCoords(t *testing.T) {
	arena := aStarTestArena{
		m:      make([]aStarTestNode, 0),
		height: 5,
		width:  5,
	}

	indexArena := [][]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24},
	}
	tests := []struct {
		x int
		y int
	}{
		{
			x: 0,
			y: 0,
		},
		{
			x: 0,
			y: 1,
		},
		{
			x: 0,
			y: 2,
		},
		{
			x: 4,
			y: 4,
		},
		{
			x: 4,
			y: 2,
		},
	}
	for i := range tests {
		test := tests[i]
		index := arena.indexFromCoords(test.x, test.y)
		require.Equal(t, indexArena[test.y][test.x], index, fmt.Sprintf("tested x: %d, y: %d", test.x, test.y))
	}
}

func TestAstar(t *testing.T) {
	data := []int{
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
	}

	arena := aStarTestArena{
		m:      make([]aStarTestNode, 0),
		height: 5,
		width:  5,
	}
	for i := range data {
		arena.m = append(arena.m, aStarTestNode{
			walkable: data[i] == 1,
			index:    i,
			cost:     1,
		})
	}
	tests := []struct {
		from int
		to   int
		len  int
	}{
		// {
		// 	from: 0,
		// 	to:   9,
		// 	len:  5,
		// },
		// {
		// 	from: 0,
		// 	to:   2,
		// 	len:  2,
		// },
		// {
		// 	from: 1,
		// 	to:   14,
		// 	len:  5,
		// },
		{
			from: 22,
			to:   4,
			len:  6,
		},
	}
	for _, test := range tests {
		res, found := Path(&arena, test.from, test.to)
		require.True(t, found)
		require.NotEmpty(t, res)
		require.Len(t, res, test.len)
	}
}
