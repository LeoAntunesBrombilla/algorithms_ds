package hash_map

import "github.com/LeoAntunesBrombilla/algorithms_ds/pkg/linked_list"

type RetailProduct struct {
	name        string
	price       float32
	stock       int
	description string
}

const arraySize = 7

type HashTable struct {
	array [arraySize]*linked_list.Slot
}

func NewHashTable() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &linked_list.Slot{}
	}
	return ht
}
