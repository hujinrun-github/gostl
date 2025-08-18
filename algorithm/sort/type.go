package sort

type SortInterface interface {
	Less(i, j int) bool // 比较函数
	Len() int           // 返回元素个数
	Swap(i, j int)      // 交换元素
}
