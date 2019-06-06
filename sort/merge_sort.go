/**
* 归并排序
* 分而治之 ，数组拆分成最小任务 然后借助一个临时数组有序合并
* @Author: zhangjian@mioji.com
* @Date: 2019/6/6 下午4:01
*/
package sort

/**
 A 数组
 start 数组开始位置
 end 数组大小
 */
func Merge_sort_c(A []int, start int, end int) {
	//递归终止
	if start >= end {
		return
	}

	//取中间位置
	mid := int((start + end) / 2)
	Merge_sort_c(A, start, mid)
	Merge_sort_c(A, mid + 1, end)
	merge(A, start, end, mid)
}

func merge(A []int, start, end, mid int) {
	var l, r = start, mid+1
	var temp []int

	for l <= mid && r <= end {
		if A[l] <= A[r] {
			temp = append(temp, A[l])
			l++
		} else {
			temp = append(temp, A[r])
			r++
		}
	}
	for mid > l {
		temp = append(temp, A[l])
		l++
	}
	for end > r {
		temp = append(temp, A[r])
		r++
	}
	for _, v := range temp {
		A[start] = v
		start++
	}
}