/**
* @Author: zhangjian@mioji.com
* @Date: 2019/6/6 下午6:22
*/
package sort

import (
	"testing"
)

func TestM(t *testing.T) {
	x := []int{7,3,234,2324,4,44,4234,44}
	Merge_sort_c(x, 0 , len(x) -1)
}