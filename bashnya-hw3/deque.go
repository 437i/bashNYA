package main

import "fmt"

type Deque struct {
	items []int
}

func DequeInit() *Deque {
	return &Deque{items: make([]int, 0)}
}

func (d *Deque) PushFront(num int) {
	d.items = append([]int{num}, d.items...)
}

func (d *Deque) PushBack(num int) {
	d.items = append(d.items, num)
}

func (d *Deque) PopFront() (int, bool) {
	if len(d.items) == 0 {
		fmt.Println("error: empty deque")
		return 0, false
	}
	first_item := d.items[0]
	d.items = d.items[1:]
	return first_item, true
}

func (d *Deque) PopBack() (int, bool) {
	if len(d.items) == 0 {
		fmt.Println("error: empty deque")
		return 0, false
	}
	last_index := len(d.items) - 1
	last_item := d.items[last_index]
	d.items = d.items[:last_index]
	return last_item, true
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) Size() int {
	return len(d.items)
}

func (d *Deque) Clear() {
	d.items = d.items[:0]
}
