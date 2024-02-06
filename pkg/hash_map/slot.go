package hash_map

import (
	"fmt"
)

type Node struct {
	key  *RetailProduct
	next *Node
}

type Slot struct {
	head *Node
}

func NewSlot() *Slot {
	return &Slot{head: nil}
}

func (b *Slot) insert(product *RetailProduct) {
	if !b.search(product.id) {
		newNode := &Node{key: product}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(product, "already exists")
	}
}

func (b *Slot) search(id string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key.id == id {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *Slot) Delete(k string) {
	if b.head.key.id == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key.id == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}

}
