package hash_map

type RetailProduct struct {
	id          string
	name        string
	price       float32
	stock       int
	description string
}

const arraySize = 7

type HashTable struct {
	array [arraySize]*Slot
}

func NewHashTable() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &Slot{}
	}
	return ht
}

func HashFunction(id string) int {
	var sum int

	for _, v := range id {
		sum += int(v)
	}
	return sum % arraySize
}

func (h *HashTable) Insert(product *RetailProduct) {
	idx := HashFunction(product.id)
	h.array[idx].insert(product)
	return
}
