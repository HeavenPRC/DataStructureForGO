/**
* @Author: zhangjian@mioji.com
* @Date: 2019/6/11 下午2:47
*/
package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	x := []int{77,  88,99, 2,10}
	QuickSort(x, 0, len(x) - 1)
	fmt.Println(x)
	x = []int{1, 2, 3, 4,10}
	QuickSort(x, 0, len(x) - 1)
	fmt.Println(x)
}