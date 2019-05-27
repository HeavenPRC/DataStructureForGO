/**
* 选择排序
* 分已排序和未排序区间，每次从未排序区间中找到最小元素 放到已排序区间末尾
* @Author: zhangjian@mioji.com
* @Date: 2019/5/27 上午10:07
*/
package sort

import "fmt"

func SelectionSort(a []int) {
	len := len(a)
	for i := 0; i < len; i++ {
		val :=  a[i]
		m := i
		for j := i + 1; j < len ;j++  {
			if val > a[j] {
				val = a[j]
				m = j
			}
		}
		a[i],a[m] = a[m],a[i]
	}
	fmt.Println(a)
}