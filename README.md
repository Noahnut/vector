[![Build Status](https://travis-ci.com/Noahnut/vector.svg?branch=main)](https://travis-ci.com/Noahnut/vector)

# Simple Vector
simple C++ vector library implemention by the golang as golang library

Use the doubly linked list to implement this vector
There serveral method like C++ vector function can be use

## data struct detail

### Size and Capacity
This Size is mean which node actually have the value in the linked list and Capacity is the how many node in the 
linked list.
If the size is bigger than the Capacity will extend the linked list Capacity to the twice of current Capacity.
## install
    go get github.com/Noahnut/vector

## Usage
```go
// Create the new Vector
v := NewVector()

// add the new element to the vector end
v.Push_back(10)

// add the new element to vector header
v.Push_front(20) 

// vector header value 20
v.Front() 

// vector end value 10
v.Back() 

// insert element to the 1 position value is 30
v.Insert(30,1)

// get the memory address of the 1 node
*v.At(1)

// take the last element out 
v.Pop_back()

// take the first element out
v.Pop_front()
```
### Vector struct
vector basic struct to record linked list header, tail and end
```go
type Vector struct {
	size     int   // element number in vector
	capacity int   // vector capacity
	header   *node // first element in vector
	tail     *node // last element in vector
	end      *node // end of the vector
}
```

### Doubly linked list node
node struct to record prev and next node
```go
type node struct {
	value interface{}
	next  *node
	prev  *node
}
```
### Create new vector
```go
func NewVector(parameter ...int) Vector
```
#### Example
1.  Default
Create the New vector without any parameter
```go
v := NewVector()
```
2. Create the New vector with size 
```go
v := NewVector(10)
```
3. Create the New vector with size and value
```go
v := NewVector(10, 10)
```

### push_back
push the element to the last of the linked list
```go
func (this *Vector) Push_back(value interface{})
```

### push_front
push the element to the front of the linked list
```go
func (this *Vector) Push_front(value interface{})
```

### pop_back
pop the last element out of the linked list
```go
func (this *Vector) Pop_back()
```

### pop_front
pop the front element out of the linked list
```go
func (this *Vector) Pop_front()
```

### Front
show first element value in the linked list 
```go
func (this *Vector) Front() interface{}
```

### Back 
show last element value in the linked list
```go
func (this *Vector) Back() interface{}
```

### Size 
show the linked list size 
```go
func (this *Vector) Size() int
```

### Capacity
show the linked list capacity
```go 
func (this *Vector) Capacity() int
```

### Empty
show the linked list is empty or not
```go 
func (this *Vector) Empty() bool
```