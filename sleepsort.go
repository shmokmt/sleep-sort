package sleepsort

import (
	"time"
)

func SortUint(s []uint) []uint {
	ch := make(chan uint)
	for _, elem := range s {
		go func(i uint, ch chan uint) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			ch <- i
		}(elem, ch)
	}
	ret := make([]uint, 0, len(s))
	for idx := 0; idx < len(s); idx++ {
		ret = append(ret, <-ch)
	}
	return ret
}
