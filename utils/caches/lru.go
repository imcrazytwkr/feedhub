package caches

type LRU[K comparable, V any] interface {
	Add(key K, value V)
	Get(key K) (value V, ok bool)
	Remove(key K)
	RemoveOldest()
	Clear()
	Len() int
}
