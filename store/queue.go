package store

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) init() {
	q.head = &Node{
		value: 0,
		next:  nil,
	}
	q.tail = q.head
	q.size = 0
}
func (q *Queue) isEmpty() bool {
	return q.size == 0
}
func (q *Queue) pop() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	v := q.head.value
	q.head = q.head.next
	q.size--
	return v, true
}
func (q *Queue) push(v int) {
	q.tail = &Node{
		value: v,
		next:  q.tail.next,
	}
}

type QueueArr []int

func (q *QueueArr) isEmpty() bool {
	return len(*q) == 0
}
func (q *QueueArr) pop() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	v := (*q)[0]
	*q = (*q)[:len(*q)-1]
	return v, true
}
func (q *QueueArr) push(v int) {
	*q = append(*q, v)
}

// 优先队列 小根堆
type PriorityQueue struct {
}

func (p *PriorityQueue) Init() {
}
func (p *PriorityQueue) Pop() {

}
func (p *PriorityQueue) Push() {

}
func (p *PriorityQueue) Delete() {

}

type BigHead struct {
}

func (bh *BigHead) Init() {

}
func (bh *BigHead) Delete() {

}
func (bh *BigHead) Add() {

}
func (bh *BigHead) Sort() {

}
