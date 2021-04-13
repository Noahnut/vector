package vector

import (
	"errors"
	"fmt"
)

/*
This is simple C++ vector implement by the golang
use the double Linked List
*/

type node struct {
	value interface{}
	next  *node
	prev  *node
}

type Vector struct {
	size     int   // element number in vector
	capacity int   // vector capacity
	header   *node // first element in vector
	tail     *node // last element in vector
	end      *node // end of the vector
}

/*
construct the vector
make the capacity as double the size which we need
*/
func (this *Vector) constructVector(size int, value interface{}) {
	firstNode := node{
		value: value,
	}
	currNode := &firstNode
	capacity := size * 2

	for i := 1; i < capacity; i++ {
		node := node{value: value}
		if value != nil && i == size-1 {
			this.tail = &node
		}
		currNode.next = &node
		node.prev = currNode
		currNode = &node
	}
	this.header = &firstNode

	if value == nil {
		this.tail = this.header
		this.size = 0
	} else {
		this.size = size
	}
	this.end = currNode

	this.capacity = capacity
}

func (this *Vector) extendSpace() {
	extendSize := this.capacity
	this.capacity = this.capacity * 2

	for i := 0; i < extendSize; i++ {
		node := node{}
		this.end.next = &node
		node.prev = this.end
		this.end = this.end.next
	}
}

func NewVector(parameter ...int) Vector {
	v := Vector{}

	parSize := len(parameter)

	switch parSize {
	case 0:
		v.constructVector(10, nil)
	case 1:
		v.constructVector(parameter[0], nil)
	case 2:
		v.constructVector(parameter[0], parameter[1])
	}
	return v
}

func (this *Vector) Push_back(value interface{}) {
	if this.tail == this.end {
		this.extendSpace()
	}
	this.tail.value = value
	if this.tail != this.end {
		this.tail = this.tail.next
	}
	this.size++
}

func (this *Vector) Push_front(value interface{}) {
	NewNode := node{}
	NewNode.value = value
	this.capacity++
	this.size++
	this.header.prev = &NewNode
	NewNode.next = this.header
	this.header = &NewNode
}

func (this *Vector) Pop_back() {
	this.tail.value = nil
	this.tail = this.tail.prev
	this.size--
}

func (this *Vector) Pop_front() {
	this.header = this.header.next
	this.size--
	this.capacity--
}

func (this *Vector) Insert(value interface{}, position int) error {
	if position > this.size {
		return errors.New("You position is out of size")
	}
	if this.size+1 > this.capacity {
		this.extendSpace()
	}

	newNode := node{value: value}

	iter := this.header
	pathCount := 0

	for pathCount != position {
		iter = iter.next
		pathCount++
	}

	if iter.prev != nil {
		iter.prev.next = &newNode
		newNode.prev = iter.prev
	}

	if iter != nil {
		iter.prev = &newNode
		newNode.next = iter
	}

	if position == 0 {
		this.header = &newNode
	}

	this.size++
	this.capacity++

	return nil
}

func (this *Vector) Erase(position int) {
	if position > this.size {
		fmt.Print("You position is out of size")
		return
	}

	iter := this.header
	pathCount := 0

	for pathCount != position {
		iter = iter.next
		pathCount++
	}

	if iter.prev != nil {
		iter.prev.next = iter.next
		if iter.next != nil {
			iter.next.prev = iter.prev
		}
	}

	if position == 0 {
		this.header = iter.next
	}

	this.size--
	this.capacity--

	iter = nil
}

func (this *Vector) Swap(v *Vector) bool {
	return false
}

func (this *Vector) Clear() {
	iter := this.header
	index := 0

	for iter != nil {
		tempNext := iter.next
		iter.value = nil
		if index >= 10 {
			iter = nil
			this.capacity--
		}
		iter = tempNext
		index++
	}
	this.size = 0
}

func (this *Vector) At(position int) *interface{} {
	iter := this.header
	index := 0
	for index != position {
		iter = iter.next
		index++
	}
	return &iter.value
}

func (this *Vector) Front() interface{} {
	return this.header.value
}

func (this *Vector) Back() interface{} {
	return this.tail.prev.value
}

func (this *Vector) Size() int {
	return this.size
}

func (this *Vector) Capacity() int {
	return this.capacity
}

func (this *Vector) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}
