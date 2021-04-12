package vector

import (
	"errors"
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

func TestAt(t *testing.T) {
	v := NewVector()

	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	if *v.At(5) != 60 {
		t.Errorf("Expect result %d result %d", 60, *v.At(5))
	}

	if *v.At(0) != 10 {
		t.Errorf("Expect result %d result %d", 10, *v.At(0))
	}

	if *v.At(30) != nil {
		t.Errorf("Expect result nil")
	}

	for i := 0; i < 200; i += 10 {
		value := (*v.At(i / 10)).(int)
		if i+10 != value {
			t.Errorf("Expect result %d result %d", i, value)
		}
	}

	*v.At(5) = 80

	if *v.At(5) != 80 {
		t.Errorf("Expect result %d result %d", 80, *v.At(5))
	}
}

func TestClear(t *testing.T) {
	v := NewVector()

	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	v.Clear()

	if v.Size() != 0 && v.Capacity() != 10 {
		t.Error("Clear Fail")
	}
}

func TestInsert(t *testing.T) {
	v := NewVector()

	for i := 10; i <= 200; i += 10 {
		v.Push_back(i)
	}

	for i := 0; i < 5; i++ {
		v.Insert(10*i, i+2)
	}

	for i := 0; i < 5; i++ {
		if *v.At(i + 2) != 10*i {
			t.Errorf("Expect result %d result %d", 10*i, *v.At(i + 2))
		}
	}

	v.Insert(70, 0)
	if v.Front() != 70 {
		t.Errorf("Expect result %d result %d", 70, v.Front())
	}

	v.Insert(90, v.Size())
	if v.Back() != 90 {
		t.Errorf("Expect result %d result %d", 70, v.Back())
	}

	if v.Capacity() != 47 {
		t.Errorf("Expect result %d result %d", 47, v.Capacity())
	}

	if v.Size() != 27 {
		t.Errorf("Expect result %d result %d", 47, v.Size())
	}

	err := v.Insert(10, 30)
	if err == errors.New("You position is out of size") {
		t.Error("The error result is not correct")
	}

	for i := 0; i <= 60; i++ {
		v.Insert(10*i, i)
	}

	if v.Capacity() != 108 {
		t.Errorf("Expect result %d result %d", 108, v.Capacity())
	}

	if v.Size() != 88 {
		t.Errorf("Expect result %d result %d", 88, v.Size())
	}
}
