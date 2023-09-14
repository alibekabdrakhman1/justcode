package main

import (
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture1/compare_slices"
	"github.com/alibekabdrakhman/justcode/lecture1/longest_common_prefix"
	"github.com/alibekabdrakhman/justcode/lecture1/sort_slice"
	"github.com/alibekabdrakhman/justcode/lecture1/two_sum"
)

func main() {

	fmt.Println("----First Task----")
	fmt.Println(two_sum.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(two_sum.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(two_sum.TwoSum([]int{3, 3}, 6))
	fmt.Println("----Second Task----")
	fmt.Println(longest_common_prefix.Prefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longest_common_prefix.Prefix([]string{"dog", "racecar", "car"}))
	fmt.Println("----Third Task----")
	fmt.Println(compare_slices.Compare([]int{1, 2, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}))
	fmt.Println(compare_slices.Compare([]int{1, 2, 3, 4, 6}, []int{4, 3, 2, 6, 1}))
	fmt.Println("----Fourth Task----")
	s1 := sort_slice.SortSlice{Slice: []int{31, 31, 42, 49, 51}}
	fmt.Println(s1.Sort())

}
