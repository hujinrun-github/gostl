package comparator

import (
	"github.com/hujinrun-github/gostl/util/generic_type"
)

type IntegerComparator[T generic_type.Integer] struct {
	sd SortDirect
}

func (i *IntegerComparator[T]) SetSortDirect(sd SortDirect) {
	i.sd = sd
}

func (i *IntegerComparator[T]) Compare(left T, right T) int {
	if (i.sd == ASCEND && left > right) || (i.sd == DESCEND && left < right) {
		return 1
	} else if (i.sd == ASCEND && left < right) || (i.sd == DESCEND && left > right) {
		return -1
	}
	return 0
}
