package comparator

type Comparator[K any] interface {
	Compare(left K, right K) int
}
