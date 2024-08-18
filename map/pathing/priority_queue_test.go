package pathing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testNode struct {
	priority int
}

func (t *testNode) GetPathCost() int {
	return t.priority
}

func (t *testNode) GetIndex() int {
	return 0
}

func (t *testNode) IsWalkable() bool {
	return true
}

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue()

	one := &testNode{priority: 5}
	queue.Push(5, one)
	two := &testNode{priority: 1}
	queue.Push(1, two)
	three := &testNode{priority: 10}
	queue.Push(10, three)
	four := &testNode{priority: 11}
	queue.Push(10, four)

	res1 := queue.Pop()
	res2 := queue.Pop()
	res3 := queue.Pop()
	res4 := queue.Pop()
	require.Equal(t, 10, res1.GetPathCost())
	require.Equal(t, 11, res2.GetPathCost())
	require.Equal(t, 5, res3.GetPathCost())
	require.Equal(t, 1, res4.GetPathCost())
}
