package hash_map

import (
	"testing"
)

func TestNewHashTable(t *testing.T) {
	ht := NewHashTable()

	if ht == nil {
		t.Error("NewHashTable() returned nil, expected a valid hash table instance")
	}

	for i, slot := range ht.array {
		if slot == nil {
			t.Errorf("Slot %d is nil, expected an initialized slot", i)
		}
	}
}

func TestHashFunction(t *testing.T) {
	product := RetailProduct{
		id:          "ADX234",
		name:        "MacBook",
		price:       1000.0,
		stock:       24,
		description: "Personal computer",
	}

	idx := HashFunction(product.id)

	if idx < 0 || idx >= arraySize {
		t.Errorf("Invalid index %d returned by HashFunction, expected index between 0 and %d", idx, arraySize-1)
	}

	if idx != 3 {
		t.Errorf("Invalid index %d returned by HashFunction, expected index 3", idx)
	}
}

func TestSlotInsertion(t *testing.T) {
	product := RetailProduct{
		id:          "ADX234",
		name:        "MacBook",
		price:       1000.0,
		stock:       24,
		description: "Personal computer",
	}

	ht := NewHashTable()

	ht.Insert(&product)

	if ht.array[3].head == nil {
		t.Error("Slot head is nil, expected a valid node")
	}

}

//func TestInsertProduct(t *testing.T) {
//	ht := NewHashTable()
//}
