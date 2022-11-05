# sleep-sort

The implementation of Sleep Sort in Go

```go
package main

import sleepsort "github.com/shmokmt/sleep-sort"

func main() {
	s := []uint{3, 5, 2, 6}
	sleepsort.SortUint(s) // [2, 3, 5, 6]
}
```

## See also

- [Sorting algorithms/Sleep sort](https://rosettacode.org/wiki/Sorting_algorithms/Sleep_sort)
- [画期的（？）なソートアルゴリズム「Sleep Sort」 | gihyo.jp](https://gihyo.jp/dev/clip/01/orangenews/vol63/0006)
