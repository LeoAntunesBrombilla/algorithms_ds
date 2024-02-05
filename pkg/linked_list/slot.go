package linked_list

import "fmt"

type Node struct {
	key  string
	next *Node
}

type Slot struct {
	head *Node
}

func NewSlot() *Slot {
	return &Slot{head: nil}
}

func (b *Slot) insert(k string) {
	if !b.search(k) {
		newNode := &Node{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

func (b *Slot) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *Slot) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}

}
