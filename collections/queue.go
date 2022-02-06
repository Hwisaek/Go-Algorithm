package collections

type Queue struct {
	queue []interface{}
}

func (q *Queue) Push(e interface{}) {
	q.queue = append(q.queue, e)
}

func (q *Queue) Pop(i interface{}) {
	q.queue = q.queue[1:]
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Front() interface{} {
	return q.queue[0]
}
func (q *Queue) Back() interface{} {
	return q.queue[q.Size()-1]
}
