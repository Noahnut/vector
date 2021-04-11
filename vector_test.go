package vector

import (
	"testing"
)

func TestConstruct(t *testing.T) {
	v := NewVector()

	if v.Empty() != true {
		t.Errorf("The vector should be empty")
	}

	if v.Capacity() != 20 {
		t.Error("defaulf capacity is wrong")
	}

	if v.Size() != 0 {
		t.Errorf("The vector should be empty")
	}

	if v.tail != v.header {
		t.Errorf("The vector should be empty")
	}

	v = NewVector(40)
	if v.Capacity() != 80 {
		t.Errorf("Expect capacity %d result %d", 80, v.capacity)
	}

	if v.Size() != 0 {
		t.Errorf("Expect size %d result %d", 40, v.size)
	}

	v = NewVector(40, 20)

	if v.Empty() != false {
		t.Errorf("The vector should be empty")
	}

	if v.Capacity() != 80 {
		t.Errorf("Expect capacity %d result %d", 80, v.capacity)
	}

	if v.Size() != 40 {
		t.Errorf("Expect size %d result %d", 40, v.size)
	}
}

func TestPushBackPopBack(t *testing.T) {
	v := NewVector()

	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	if v.Size() != 20 {
		t.Errorf("Expect size %d result %d", 20, v.Size())
	}

	if v.Capacity() != 40 {
		t.Errorf("Expect size %d result %d", 40, v.Capacity())
	}

	if v.Back() != 200 {
		t.Errorf("Expect result %d result %d", 200, v.Back())
	}

	for i := 210; i <= 500; i += 10 {
		v.Push_back(i)
	}

	if v.Size() != 50 {
		t.Errorf("Expect size %d result %d", 50, v.Size())
	}

	if v.Capacity() != 80 {
		t.Errorf("Expect size %d result %d", 80, v.Capacity())
	}

	if v.Back() != 500 {
		t.Errorf("Expect result %d result %d", 500, v.Back())
	}
}

func TestPushFront(t *testing.T) {
	v := NewVector()

	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	if v.Front() != 10 {
		t.Errorf("Expect result %d result %d", 10, v.Front())
	}

	v.Push_front(20)

	if v.Size() != 21 {
		t.Errorf("Expect result %d result %d", 21, v.Size())
	}

	if v.Capacity() != 41 {
		t.Errorf("Expect result %d result %d", 41, v.Capacity())
	}

	if v.Front() != 20 {
		t.Errorf("Expect result %d result %d", 20, v.Front())
	}

	for i := 210; i <= 400; i += 10 {
		v.Push_front(i)
	}

	if v.Size() != 41 {
		t.Errorf("Expect result %d result %d", 21, v.Size())
	}

	if v.Capacity() != 61 {
		t.Errorf("Expect result %d result %d", 41, v.Capacity())
	}

	if v.Front() != 400 {
		t.Errorf("Expect result %d result %d", 20, v.Front())
	}
}

func TestPopFront(t *testing.T) {
	v := NewVector()
	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	if v.Size() != 20 {
		t.Errorf("Expect result %d result %d", 20, v.Size())
	}

	if v.Capacity() != 40 {
		t.Errorf("Expect result %d result %d", 40, v.Capacity())
	}

	if v.Front() != 10 {
		t.Errorf("Expect result %d result %d", 10, v.Front())
	}

	v.Pop_front()

	if v.Size() != 19 {
		t.Errorf("Expect result %d result %d", 19, v.Size())
	}

	if v.Capacity() != 39 {
		t.Errorf("Expect result %d result %d", 39, v.Capacity())
	}

	if v.Front() != 20 {
		t.Errorf("Expect result %d result %d", 20, v.Front())
	}

	for i := 210; i <= 400; i += 10 {
		v.Push_front(i)
	}

	if v.Size() != 39 {
		t.Errorf("Expect result %d result %d", 39, v.Size())
	}

	if v.Capacity() != 59 {
		t.Errorf("Expect result %d result %d", 59, v.Capacity())
	}

	if v.Front() != 400 {
		t.Errorf("Expect result %d result %d", 400, v.Front())
	}

	for i := 0; i < 10; i++ {
		v.Pop_front()
	}

	if v.Size() != 29 {
		t.Errorf("Expect result %d result %d", 29, v.Size())
	}

	if v.Capacity() != 49 {
		t.Errorf("Expect result %d result %d", 49, v.Capacity())
	}

	if v.Front() != 300 {
		t.Errorf("Expect result %d result %d", 300, v.Front())
	}
}

func TestPopBack(t *testing.T) {
	v := NewVector()

	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	v.Pop_back()
	v.Pop_back()
	v.Pop_back()
	v.Pop_back()

	if v.Size() != 16 {
		t.Errorf("Expect result %d result %d", 16, v.Size())
	}

	if v.Capacity() != 40 {
		t.Errorf("Expect result %d result %d", 40, v.Capacity())
	}

	if v.Back() != 160 {
		t.Errorf("Expect result %d result %d", 160, v.Back())
	}

	for i := 160; i <= 460; i += 10 {
		v.Push_back(i)
	}

	v.Pop_back()
	v.Pop_back()
	v.Pop_back()
	v.Pop_back()

	if v.Size() != 43 {
		t.Errorf("Expect result %d result %d", 43, v.Size())
	}

	if v.Capacity() != 80 {
		t.Errorf("Expect result %d result %d", 80, v.Capacity())
	}

	if v.Back() != 420 {
		t.Errorf("Expect result %d result %d", 420, v.Back())
	}
}
