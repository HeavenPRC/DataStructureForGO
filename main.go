/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/17 下午2:39
*/
package main

import (
	x "DataStructureForGO/sort"
	"fmt"
)

func main() {
	A := [...]int{9,3,4,7,1}
    x.Merge_sort_c(A[:], 0, 4)
	fmt.Println(A)
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}