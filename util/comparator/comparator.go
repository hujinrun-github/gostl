package comparator

type SortDirect int

const (
	ASCEND SortDirect = iota
	DESCEND
)

type Comparator[K any] interface {
	Compare(left K, right K) int
	SetSortDirect(sd SortDirect)
}
