package pathing

import "container/heap"

type INode interface {
	GetPathCost() int
	GetIndex() int
	IsWalkable() bool
}

func newPriorityQueueInternal() *priorityQueueInternal {
	return &priorityQueueInternal{
		Content: make([]nodeWithPriority, 0),
	}
}

type nodeWithPriority struct {
	priority int
	node     INode
}

type priorityQueueInternal struct {
	Content []nodeWithPriority
}

func (p *priorityQueueInternal) Push(x any) { // add x as element Len()
	node := x.(nodeWithPriority)
	p.Content = append(p.Content, node)
}

func (p *priorityQueueInternal) Pop() any { // remove and return element Len() - 1.
	pop := p.Content[len(p.Content)-1]
	p.Content = p.Content[0 : len(p.Content)-1]
	return pop
}

func (p *priorityQueueInternal) Len() int {
	return len(p.Content)
}

func (p *priorityQueueInternal) Less(i, j int) bool {
	return p.Content[i].priority < p.Content[j].priority
}

func (p *priorityQueueInternal) Swap(i, j int) {
	p.Content[i], p.Content[j] = p.Content[j], p.Content[i]
}

type PriorityQueue struct {
	internalQueue *priorityQueueInternal
}

func NewPriorityQueue() *PriorityQueue {
	q := &PriorityQueue{
		internalQueue: newPriorityQueueInternal(),
	}
	heap.Init(q.internalQueue)
	return q
}

func (p *PriorityQueue) Push(priority int, node INode) {
	heap.Push(p.internalQueue, nodeWithPriority{priority: priority, node: node})
}

func (p *PriorityQueue) Pop() INode {
	popped := heap.Pop(p.internalQueue).(nodeWithPriority)
	return popped.node
}

func (p *PriorityQueue) Len() int {
	return p.internalQueue.Len()
}
