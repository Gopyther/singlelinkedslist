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
	return &PowerNode{PNodeMessage: "This is a message from the power Node"}
}

// func main() {
// 	var n Node
// 	n = &SLLNode{}
// 	fmt.Println(n.SetValue(4))
// }

// func main() {
// 	printType("text")
// 	printType(3)
// 	printType(4.0)
// }

// func printType(i interface{}) {
// 	switch i := i.(type) {
// 	case string:
// 		fmt.Println("This is a string type", i)
// 	case int:
// 		fmt.Println("This is an int type", i)
// 	case float32:
// 		fmt.Println("This is a float type", i)
// 	}
// }
// func main() {
// 	n := createNode(5)

// 	switch concreten := n.(type) {
// 	case *SLLNode:
// 		fmt.Println("Type is SLLNode, message:", concreten.SNodeMessage)
// 	case *PowerNode:
// 		fmt.Println("Type is PowerNode, message:", concreten.PNodeMessage)
// 	}
// 	sNode := &SLLNode{}
// 	n = sNode
// }

func main() {
	n := createNode(5)

	switch concreten := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode, message:", concreten.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message:", concreten.PNodeMessage)
	}
	sNode := &SLLNode{value: 15}
	fmt.Println(sNode.GetValue())
}
func createNode(v int) Node {
	sn := NewSLLNode()
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

// type magicStore struct {
// 	value interface{}
// 	name  string
// }

// func (ms *magicStore) SetValue(v interface{}) {
// 	ms.value = v
// }

// func (ms *magicStore) GetValue() interface{} {
// 	return ms.value
// }

// func NewMagicStore(nm string) *magicStore {
// 	return &magicStore{name: nm}
// }

// func main() {
// 	istore := NewMagicStore("Integer Store")
// 	istore.SetValue(4)
// 	if v, ok := istore.GetValue().(int); ok {
// 		v *= 4
// 		fmt.Println(v)
// 	}

// 	sstore := NewMagicStore("String store")
// 	sstore.SetValue("Hello")
// 	if v, ok := sstore.GetValue().(string); ok {
// 		v += " World"
// 		fmt.Println(v)
// 	}

// }
