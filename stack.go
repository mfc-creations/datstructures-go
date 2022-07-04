package main

type Stack struct {
	item []int64
}

func (s *Stack) Push(num int64) {
	s.item = append(s.item, num)

}

func (s *Stack) Pop() {
	s.item = s.item[:len(s.item)-1]
}

// func main() {
// 	fmt.Print("Stack")
// 	myStack := Stack{}
// 	fmt.Print("Before", myStack.item)
// 	myStack.Push(1)
// 	myStack.Push(13)
// 	fmt.Print("After", myStack.item)
// 	myStack.Pop()
// 	fmt.Print("After", myStack.item)

// }
