package queue

type Squeue []interface{}

func New() *Squeue {
	q := make(Squeue, 0)
	return &q
}

func (q *Squeue) Push(i interface{}) {
	*q = append(*q, i)
}

func (q *Squeue) Pop() (interface{}, bool) {
	if len(*q) == 0 {
		return 0, false
	}

	i := (*q)[0]
	*q = (*q)[1:]

	if n := cap(*q) / 2; len(*q) <= n {
		nodes := make([]interface{}, len(*q), n)
		copy(nodes, *q)
		*q = nodes
	}

	return i, true

}

func (q *Squeue) Cap() int {
	return cap(*q)
}

func (q *Squeue) Len() int {
	return len(*q)
}
