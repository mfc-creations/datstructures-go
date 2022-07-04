package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(num int) {
	q.items = append(q.items, num)
}

func (q *Queue) Dequeue() {
	q.items = q.items[1:]
}

// func main() {
// 	myQueue := Queue{}
// 	fmt.Print("Before", myQueue.items)
// 	myQueue.Enqueue(1)
// 	myQueue.Enqueue(2)
// 	fmt.Print("After enqueue", myQueue.items)
// 	myQueue.Dequeue()
// 	fmt.Print("After dequeue", myQueue.items)

// }
