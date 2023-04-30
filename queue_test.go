package queue

import (
	"testing"
)

func Test_PushBack(t *testing.T) {
	in := []int{1, 12, 3, 4568, 9, 10, -34}

	queue := Queue[int]{}
	for _, v := range in {
		queue.PushBack(v)
	}

	for i := 0; i < len(in); i++ {
		if in[i] != queue[i] {
			t.Errorf("queue.push(%d)=>%d;want %d", in[i], queue[i], in[i])
		}
	}
}

func Test_PopFront(t *testing.T) {
	in := []int{1, 12, 3, 4568, 9, 10, -34}
	queue := make(Queue[int], len(in))
	copy(queue, in)

	for i := 0; i < len(in); i++ {
		if out, err := queue.PopFront(); err != nil || out != in[i] {
			t.Errorf("out:%d want:%d", out, in[i])
		}
	}
}

func Test_PopBack(t *testing.T) {
	in := []int{1, 12, 3, 4568, 9, 10, -34}
	queue := make(Queue[int], len(in))
	copy(queue, in)

	for i := len(in) - 1; i >= 0; i-- {
		if out, err := queue.PopBack(); err != nil || out != in[i] {
			t.Errorf("out:%d want:%d", out, in[i])
		}
	}
}
