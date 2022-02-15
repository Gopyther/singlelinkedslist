package main

import (
	"errors"
	"fmt"
)

var ErrInvalidNode = errors.New("Node is not valid")

type Node interface {
	SetValue(v int) error
	GetValue() int
}

//type SLLNode
type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v
	return nil
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is a message from the normal Node"}
}

//type PowerNode
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

func (sNode *PowerNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v * 10
	return nil
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is a message from the normal Node"}
}

func main() {

	n := createNode(5)

	switch concreten := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode, message:", concreten.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message:", concreten.PNodeMessage)
	}

}

func createNode(v int) Node {
	sn := NewPowerNode()
	sn.SetValue(v)
	return sn
}

// type SingleLinkedList struct {
// 	head *SLLNode
// 	tail *SLLNode
// }

// func newSingleLinkedList() *SingleLinkedList {
// 	return new(SingleLinkedList)
// }

// func (list *SingleLinkedList) Add(v int) {
// 	newNode := &SLLNode{value: v}
// 	if list.head == nil {
// 		list.head = newNode
// 	} else if list.tail == list.head {
// 		list.head.next = newNode
// 	} else if list.tail != nil {
// 		list.tail.next = newNode
// 	}
// 	list.tail = newNode
// }

// func (list *SingleLinkedList) String() string {
// 	s := ""
// 	for n := list.head; n != nil; n = n.next {
// 		s += fmt.Sprintf("{%d} ", n.GetValue())
// 	}
// 	return s
// }
