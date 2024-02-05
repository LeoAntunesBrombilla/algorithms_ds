package linked_list

import "testing"

func TestInsert(t *testing.T) {
	slot := NewSlot()
	key := "testKey"

	slot.insert(key)
	if !slot.search(key) {
		t.Errorf("key %s was not found after insert", key)
	}

	slot.insert(key)
	//TODO Verify the duplicate doesn't create a second node, but this requires traversing the list or modifying functionality to confirm.
}

func TestSearch(t *testing.T) {
	slot := NewSlot()
	key := "testKey"
	nonExistentKey := "nonExistentKey"

	slot.insert(key)
	if !slot.search(key) {
		t.Errorf("key %s was not found by search, but it should have been", key)
	}

	if slot.search(nonExistentKey) {
		t.Errorf("search falsely found non-existent key %s", nonExistentKey)
	}
}

func TestDelete(t *testing.T) {
	slot := NewSlot()
	key := "testKey"

	slot.insert(key)
	slot.delete(key)
	if slot.search(key) {
		t.Errorf("key %s was found after delete, but it should not have been", key)
	}

	slot.delete("someRandomKey")
}
