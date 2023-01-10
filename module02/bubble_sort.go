package module02

import (
	"sort"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for i := 0; i < len(list)-1; i++ {
		done := true
		for j := 1; j < len(list)-i; j++ {
			a, b := list[j-1], list[j]
			if a > b {
				done = false
				list[j-1], list[j] = b, a
			}
		}
		if done {
			break
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	for i := 0; i < len(list)-1; i++ {
		done := true
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
				done = false
			}
		}
		if done {
			break
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
	for i := 0; i < len(people)-1; i++ {
		done := true
		for j := 1; j < len(people)-i; j++ {
			a, b := people[j-1], people[j]
			if a.Age < b.Age {
				continue
			}
			if a.Age == b.Age && a.LastName < b.LastName {
				continue
			}
			if a.Age == b.Age && a.LastName == b.LastName && a.FirstName < b.FirstName {
				continue
			}
			done = false
			people[j-1], people[j] = b, a
		}
		if done {
			break
		}
	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	// normally we would just use:
	// sort.Sort(list)
	for i := 0; i < list.Len()-1; i++ {
		for j := 1; j < list.Len()-i; j++ {
			if list.Less(j, j-1) {
				list.Swap(j, j-1)
			}
		}
	}
}
