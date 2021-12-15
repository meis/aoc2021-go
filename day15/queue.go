package main

type queue struct {
	present map[point]bool
	sorted  []point
}

func (q *queue) empty() bool {
	return len(q.present) == 0
}
func (q *queue) next() point {
	value, list := q.sorted[0], q.sorted[1:]
	q.sorted = list
	delete(q.present, value)
	return value
}
func (q *queue) append(p point) {
	if !q.present[p] {
		q.sorted = append(q.sorted, p)
		q.present[p] = true
	}
}
