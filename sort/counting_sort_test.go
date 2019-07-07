/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/7 下午5:19
*/
package sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	fmt.Println("测试桶排序")
	a := []int{7, 10, 15,4,7,1,0,1}
	CountingSort(a)
	fmt.Println(a)
	fmt.Println("桶排序测试结束")
}