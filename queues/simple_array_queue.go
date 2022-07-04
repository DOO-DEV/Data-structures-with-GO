package queues

import "fmt"

type ArrayQueue struct {
	front    int
	rear     int
	size     int
	capacity int
	queue    []int
}

func (q *ArrayQueue) IsFull() bool {
	if q.capacity == q.size {
		fmt.Println("Queue is full")
		return true
	}
	return false
}

func (q *ArrayQueue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *ArrayQueue) EnQueue(item int) {
	if q.IsFull() {
		fmt.Println("Queue is overflwing")
		return
	}
	q.size++
	q.queue[q.rear] = item
	q.rear++
}

func (q *ArrayQueue) DeQueue() {
	if q.IsEmpty() {
		fmt.Println("Queue is underflowing")
		return
	}
	q.size--
	q.front--
	q.queue = q.queue[1:]
}

func (q *ArrayQueue) Front() int {
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
		return -1
	}
	return q.queue[q.front]
}

func (q *ArrayQueue) Rear() int {
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
		return -1
	}
	return q.queue[q.rear]
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) PrintQueue() {
	for _, val := range q.queue {
		fmt.Println(val)
	}
}

func SimpleQueueWithArray() {
	queue := ArrayQueue{capacity: 5, front: 0, rear: 0, size: 0, queue: make([]int, 5)}
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.PrintQueue()
	queue.DeQueue()
	queue.DeQueue()
	fmt.Println("After Dequeue")
	queue.PrintQueue()
	fmt.Printf("Size of the Queue is %d", queue.Size())

}
