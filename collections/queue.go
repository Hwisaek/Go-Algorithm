package collections

type Queue struct {
	Queue []int
}

func NewQueue() *Queue {
	return &Queue{Queue: []int{}}
}

func (q *Queue) Push(e int) {
	q.Queue = append(q.Queue, e)
}

func (q *Queue) Pop(idx int) int {
	e := q.Queue[idx]
	q.Queue = append(q.Queue[:idx], q.Queue[idx+1:]...)
	return e
}

func (q *Queue) Size() int {
	return len(q.Queue)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Front() int {
	return q.Queue[0]
}
func (q *Queue) Back() int {
	return q.Queue[q.Size()-1]
}
