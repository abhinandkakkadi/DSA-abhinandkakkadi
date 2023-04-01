package main
const size = 7

type HashTable struct {
	array [size]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}

func Init() *HashTable {

	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}


func main() {

	table := Init()
	table.Insert("Abhinand")

	

}

func (h *HashTable) Insert(k string) {

	index := hash(k)
	h.array[index].insert(k)

}

func (b *bucket) insert(k string) {

	node := &bucketNode{k,nil}
	if b.head == nil {
		b.head = node
		return
	} else {
		node.next = b.head
		b.head = node
		return
	}

}

// hash function
func hash(k string ) int {

	sum := 0
	for _,val := range k {
		sum += int(val)
	}
	return sum


}

