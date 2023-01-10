package module02

import (
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	new := []int{}
	for _, n := range list {
		ins := false
		for i, el := range new {
			if n <= el {
				new = append(new[:i+1], new[i:]...)
				new[i] = n
				ins = true
				break
			}
		}
		if !ins {
			new = append(new, n)
		}
	}
	for i, n := range new {
		list[i] = n
	}
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
// NOTE: implementing sort.Interface instead of just copy-pasting previous
// solution
func InsertionSortString(list []string) {
	sl := StringList(list)
	sort.Sort(sl)
}

type StringList []string

func (list StringList) Less(i, j int) bool {
	return list[i] <= list[j]
}
func (list StringList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
func (list StringList) Len() int {
	return len(list)
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	sort.Sort(list)
}
