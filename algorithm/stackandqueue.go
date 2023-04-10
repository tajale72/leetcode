package algorithm

import "log"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() {
	if len(*s) == 0 {
		log.Println("empty stack")
	}

	*s = (*s)[:len(*s)-1]
}

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue(val int) {
	if len(*q) == 0 {
		log.Println("empty stack")
	}

	*q = (*q)[len(*q)-1:]
}
