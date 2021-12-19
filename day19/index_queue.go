package main

type IndexQueue struct {
	sorted []int
}

func NewIndexQueue() IndexQueue {
	return IndexQueue{[]int{}}
}

func (q *IndexQueue) empty() bool {
	return len(q.sorted) == 0
}
func (q *IndexQueue) next() int {
	value, list := q.sorted[0], q.sorted[1:]
	q.sorted = list
	return value
}
func (q *IndexQueue) add(i int) {
	q.sorted = append(q.sorted, i)
}
