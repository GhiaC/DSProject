package Order

import (
	"DS/Arraylist"
	"fmt"
)

// store orders with arraylist
type TreeOrder struct {
	Orders *arraylist.List
}

type Order struct {
	Customer string
	Priority int
}

// return an instance of maxheap
func New() *TreeOrder {
	return &TreeOrder{Orders: arraylist.New()}
}
func (field *TreeOrder) List(){
	for i:= 0;i< field.Size() ;i++  {
		fmt.Println(field.getCurrentElement(i))
	}
}

func (field *TreeOrder) RightChildIndex(i int) int {
	return 2*i + 2
}

func (field *TreeOrder) ParentIndex(i int) int {
	return (i - 1) / 2
}

func (field *TreeOrder) LeftChildIndex(i int) int {
	return 2*i + 1
}

func (field *TreeOrder) getParentElement(i int) Order {
	return field.getCurrentElement(field.ParentIndex(i))
}

func (field *TreeOrder) getCurrentElement(i int) Order {
	elementCurrent, _ := field.Orders.Get(i)
	return elementCurrent.(Order)
}

func (field *TreeOrder) traceUp(i int) {
	for i > 0 && field.getParentElement(i).Priority < field.getCurrentElement(i).Priority {
		field.Orders.Swap(i,field.ParentIndex(i))
		i = field.ParentIndex(i)
	}
}

func (field *TreeOrder) traceDown(i int) {
	index := i
	leftNum := field.LeftChildIndex(i)
	if leftNum < field.Size()-1 && field.getCurrentElement(leftNum).Priority > field.getCurrentElement(index).Priority{
		index = leftNum
	}
	rightNum := field.RightChildIndex(i)
	if rightNum < field.Size()-1 && field.getCurrentElement(rightNum).Priority > field.getCurrentElement(index).Priority {
		index = rightNum
	}
	if i != index {
		field.Orders.Swap(i,index)
		field.traceDown(index)
	}
}

// insert order to arraylist of Maxheap struct
func (field *TreeOrder) Add(item Order) {
	field.Orders.Add(item)
	field.traceUp(field.Size() - 1)
}

// dequeue element with biggest priority

func (field *TreeOrder) DeQueue() Order {
	first := field.getCurrentElement(0)
	lastChild := field.Size() - 1
	field.Orders.Swap(0,lastChild)
	field.Orders.Remove(lastChild)
	field.traceDown(0)
	return first
}

// size of arraylist order
func (field *TreeOrder) Size() int {
	return field.Orders.Size()
}
