package pathing

import (
	"slices"
)

type Map interface {
	GetApproximateDistance(from, to int) float64
	GetNodeAtIndex(idx int) INode
	GetNeighbors(idx int) []INode
}

func Path(arena Map, fromIdx, toIdx int) ([]INode, bool) {
	queue := NewPriorityQueue()
	startNode := arena.GetNodeAtIndex(fromIdx)

	queue.Push(0, startNode)

	cameFrom := make(map[int]int)
	costSoFar := make(map[int]int)
	cameFrom[fromIdx] = -1
	costSoFar[fromIdx] = 0

	path := []INode{}
	for {
		if queue.Len() == 0 {
			return nil, false // not found
		}

		current := queue.Pop()

		if current.GetIndex() == toIdx {
			for {
				if cameFrom[current.GetIndex()] == -1 {
					path = append(path, current)
					slices.Reverse(path)
					return path, true
				}
				path = append(path, current)
				current = arena.GetNodeAtIndex(cameFrom[current.GetIndex()])
			}
		}

		for _, next := range arena.GetNeighbors(current.GetIndex()) {
			if !next.IsWalkable() {
				continue
			}
			newCost := costSoFar[current.GetIndex()] + next.GetPathCost()
			oldCostForNext, ok := costSoFar[next.GetIndex()]
			if !ok || newCost < oldCostForNext {
				costSoFar[next.GetIndex()] = newCost
				priority := newCost + int(arena.GetApproximateDistance(next.GetIndex(), toIdx))

				queue.Push(priority, next)
				cameFrom[next.GetIndex()] = current.GetIndex()
			}
		}
	}
}
