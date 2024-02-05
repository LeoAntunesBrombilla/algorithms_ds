package hash_map

import "testing"

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

func TestSlotInsertion(t *testing.T) {

}

//func TestInsertProduct(t *testing.T) {
//	ht := NewHashTable()
//}
