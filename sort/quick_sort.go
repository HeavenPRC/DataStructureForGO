/**
* 快速排序
* 分而治之 ，选择一个中间元素 大于小于这个数的分别放在两边
* 递归公式 ：quick_sort(p..r) = quick(p...q-1) + q + quick(q + 1...r)//p r是下标
* 终止条件：p >= r 即开始下标大于等于最大的下标
* @Author: zhangjian@mioji.com
* @Date: 2019/6/10 上午11:35
*/
package sort

import "fmt"

func QuickSort(A []int, s, e int) {
	if s >= e {
		return
	}
	pivot := A[e]
	m, i , j:= e, s, s
	for ;j <= e; j++ {
		if A[j] <= pivot {
			A[i],A[j] = A[j], A[i]
			if j == e {
				m = i
			}
			i++
		}
	}
	fmt.Println(A, A[s:e], i-1)
	QuickSort(A, s, m -1)
	QuickSort(A, m + 1, e)
}