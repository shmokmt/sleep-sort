package sleepsort_test

import (
	"reflect"
	"testing"

	sleepsort "github.com/shmokmt/sleep-sort"
)

func TestSortUint(t *testing.T) {
	input := []uint{3, 5, 2, 6}
	got := sleepsort.SortUint(input)
	want := []uint{2, 3, 5, 6}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
