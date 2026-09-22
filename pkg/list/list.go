package list

import (
	"fmt"
	"strconv"
)

// Define Str for List generic implementation
type Str string

func (s Str) String() string {
	return string(s)
}

// Define Int for List generic implementaiton

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// List struct.
type List[T fmt.Stringer] struct {
	head *Link[T]
	tail *Link[T]
}

// Create a new list.
func NewList[T fmt.Stringer]() *List[T] {
	return &List[T]{nil, nil}
}

// Get a pointer to the head of the list.
func (list *List[T]) PeekHead() *Link[T] {
	return list.head
}

// Get a pointer to the tail of the list.
func (list *List[T]) PeekTail() *Link[T] {
	return list.tail
}

// Add an element to the start of the list. Returns the added link.
func (list *List[T]) PushHead(value T) *Link[T] {
	if list.head == nil { //empty list
		newLink := &Link[T]{list, nil, nil, value}
		list.head = newLink
		list.tail = newLink
	} else { //already has a head item
		tempHead := list.head //old head lol
		//prev, next
		newLink := &Link[T]{list, nil, tempHead, value}
		tempHead.prev = newLink
		list.head = newLink
	}
	return list.head
}

// Add an element to the end of the list. Returns the added link.
func (list *List[T]) PushTail(value T) *Link[T] {
	if list.tail == nil { //if doesn't have a tail
		newLink := &Link[T]{list, nil, nil, value}
		list.head = newLink
		list.tail = newLink
	} else {
		tempTail := list.tail //value
		//prev, next
		newLink := &Link[T]{list, tempTail, nil, value}
		tempTail.next = newLink
		list.tail = newLink
	}
	return list.tail
}

// Find an element in a list given a boolean function, f, that evaluates to true on the desired element.
func (list *List[T]) Find(f func(*Link[T]) bool) *Link[T] {
	if list.head == nil { //empty list
		//var nothing T // have to initialize as type t instead of nil to account for all types
		return nil//&Link[T]{nil, nil, nil, nothing}
	}

	temp := list.head
	for {
		if f(temp) == true {
			return temp
		} else if temp == list.tail { //if at the end of the list and item not found
			return nil
		}
		temp = temp.next
	}
}

// Apply a function to every element in the list.
func (list *List[T]) Map(f func(*Link[T])) {
	if list.head != nil { //list has something in it
		temp := list.head
		for { 
			f(temp) //no return item for f so can't assign back
			if temp == list.tail{
				break
			} else { 
				temp = temp.next
			}
		}
	}
}
/////////////////////////////////////////////////////////////////////////////////////////
// Link struct.
type Link[T fmt.Stringer] struct {
	list  *List[T]
	prev  *Link[T]
	next  *Link[T]
	value T
}

// Get the list that this link is a part of.
func (link *Link[T]) GetList() *List[T] {
	return link.list
}

// Get the link's value.
func (link *Link[T]) GetValue() T {
	return link.value
}

// Set the link's value.
func (link *Link[T]) SetValue(value T) {
	link.value = value
}

// Get the link's prev.
func (link *Link[T]) GetPrev() *Link[T] {
	return link.prev
}

// Get the link's next.
func (link *Link[T]) GetNext() *Link[T] {
	return link.next
}

// Remove the link that calls PopSelf() from its list.
func (link *Link[T]) PopSelf() {
	prevLink := link.prev //GetPrev()
	nextLink := link.next //GetNext()

	if prevLink == nil { //new head
		if nextLink == nil { //list now empty
			link.list.head = nil
			link.list.tail = nil
		} else {
			nextLink.prev = nil
			link.list.head = nextLink
		}
	} else if nextLink == nil { //new tail
		prevLink.next = nil
		link.list.tail = prevLink //have to update the tail of the list 
	} else { //surrounded in list
		prevLink.next = nextLink
	    nextLink.prev = prevLink
	}
	
	link = nil //signal removal
}
