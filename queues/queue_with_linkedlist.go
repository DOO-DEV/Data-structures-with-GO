package queues

import "fmt"

type Node struct {
	data int
	next *Node
}

type QueueWithLinkedList struct {
	front *Node
	rear  *Node
}

func (q *QueueWithLinkedList) IsEmpty() bool {
	return q.front.data == 0
}

func (q *QueueWithLinkedList) EnQueue(item int) {
	newNode := Node{item, nil}
	if q.rear == nil {
		q.rear = &newNode
		q.front = &newNode
		return
	}
	q.rear.next = &newNode
	q.rear = &newNode
}

func (q *QueueWithLinkedList) DeQueue() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return 0
	}
	temp := q.front
	q.front = temp.next
	if q.front == nil {
		q.rear = nil
	}
	return temp.data
}

func (q *QueueWithLinkedList) PrintQueue() {
	curr := q.front
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func SimpleQueueWithLinkeList() {
	queue := QueueWithLinkedList{}
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	fmt.Println("Before Dequeueing")
	queue.PrintQueue()
	queue.DeQueue()
	queue.DeQueue()
	fmt.Println("After Dequeueing")
	queue.PrintQueue()
}
