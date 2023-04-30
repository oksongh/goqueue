package queue

import (
	"errors"
	"fmt"
)

// Queue is extension of slice as FIFO and FILO
//	index  0,1,2,3,4,5
// 	front [0,2,4,5,3,5] back
type Queue[T any] []T

var (
	// empty error
	ErrEmpty = errors.New("empty")
)

func (q *Queue[T]) PushBack(v ...T) {
	*q = append(*q, v...)
}

// PopFront pop from FIFO,Queue.
//
// if len(queue) == 0 ,return (queue[0],`nil`) and shorten queue.
// if not return (new(T), empty error)
func (q *Queue[T]) PopFront() (T, error) {
	if len(*q) == 0 {
		var v T
		return v, fmt.Errorf("Queue: %w", ErrEmpty)
	}
	defer func() { *q = (*q)[1:] }()
	return (*q)[0], nil
}

// PopBack pop from FILO,stack.
//
// if len(queue) == 0, return (last element,`nil`) and shorten queue.
// if not return (new(T), `empty` error)
func (q *Queue[T]) PopBack() (T, error) {
	if len(*q) == 0 {
		var v T
		return v, fmt.Errorf("Queue: %w", ErrEmpty)
	}
	n := len(*q) - 1
	defer func() { *q = (*q)[:n] }()
	return (*q)[n], nil
}

// Remove remove q[i] and shorten queue
func (q *Queue[T]) Remove(i int) error {
	if len(*q) == 0 {
		return fmt.Errorf("Queue: %w", ErrEmpty)
	}
	*q = append((*q)[0:i], (*q)[i+1:]...)
	return nil
}

// Last return last element
// if Queue is empty,return `empty` error
func (q *Queue[T]) Last() (T, error) {
	len := len(*q)
	if len == 0 {
		var v T
		return v, fmt.Errorf("Queue: %w", ErrEmpty)
	}
	return (*q)[len-1], nil
}
