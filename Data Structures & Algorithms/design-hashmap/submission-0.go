const bucketCount = 100_000

type kv struct {
	key   int
	value int
}

type MyHashMap struct {
	buckets []list.List // linkedlist
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([]list.List, bucketCount),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % bucketCount
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	bucket := &this.buckets[idx]

	for e := bucket.Front(); e != nil; e = e.Next() {
		if kv, ok := e.Value.(kv); ok && kv.key == key {
			kv.value = value
			e.Value = kv
			return
		}
	}

	bucket.PushBack(kv{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	idx := this.hash(key)
	bucket := &this.buckets[idx]

	for e := bucket.Front(); e != nil; e = e.Next() {
		if kv, ok := e.Value.(kv); ok && kv.key == key {
			return kv.value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := this.hash(key)
	bucket := &this.buckets[idx]

	for e := bucket.Front(); e != nil; e = e.Next() {
		if kv, ok := e.Value.(kv); ok && kv.key == key {
			bucket.Remove(e)
			return
		}
	}
}