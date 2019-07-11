/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/8 下午12:05
*/
package search

//非递归实现
func Bsearch(a []int, searchNum int) int {
	low := 0
	high := len(a) - 1
	//防止提前跳出循环
	for low <= high {
		//用+法可能会导致数据溢出 位运算会比
		mid := low + ((high - low) >> 1)

		if searchNum == a[mid] {
			return mid
		} else if searchNum < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//递归实现
func Bsearch2(a []int, searchNum int) int {
	return search(a, searchNum, 0, len(a) - 1)
}

func search(a []int, searchdate, low, high int) int {
	mid := low + ((high - low) >> 1)

	if searchdate == a[mid] {
		return mid
	} else if searchdate < a[mid] {
		return search(a, searchdate, low, mid -1)
	} else {
		return search(a, searchdate, mid + 1, high)
	}
}

//二分变形 查找第一个符合的值的元素
func BsearchFirst(a []int, searchNum int) int {
	low := 0
	high := len(a) - 1
	//防止提前跳出循环
	for low <= high {
		//用+法可能会导致数据溢出 位运算会比
		mid := low + ((high - low) >> 1)

		if searchNum == a[mid] {
			if (mid == 0) || (a[mid-1] != searchNum) {
				return mid
			} else {
				high = mid - 1
			}
		} else if searchNum < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//查找最后一个符合要求的元素
func BsearchEnd(a []int, searchNum int) int {
	low := 0
	high := len(a) - 1
	//防止提前跳出循环
	for low <= high {
		//用+法可能会导致数据溢出 位运算会比
		mid := low + ((high - low) >> 1)

		if searchNum == a[mid] {
			if (mid == 0) || (a[mid+1] != searchNum) {
				return mid
			} else {
				low = mid + 1
			}
		} else if searchNum < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//查找大于给定值的元素
func BsearchEqFirst(a []int, searchNum int) int {
	low := 0
	high := len(a) - 1
	//防止提前跳出循环
	for low <= high {
		//用+法可能会导致数据溢出 位运算会比
		mid := low + ((high - low) >> 1)
		if searchNum > a[mid] {
			if (mid == 0) || (a[mid-1] < searchNum) {
				return mid
			} else {
				low = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

//搜索旋转排序数组
//思路：先找到旋转点(即最小值),再利用二分查找
func SearchAb(a []int) int {
	//先找到旋转点 复杂度logn
	lens := len(a)
	low := 0
	high := lens - 1

	for low <= high {
		pivot := low + (high - low) / 2

		if a[pivot] < a[0] {
			if pivot == 0 || a[pivot - 1] >= a[0] {
				return pivot
			} else {
				high = pivot - 1
			}
		} else {
			low = pivot + 1
		}
	}
	return -1
}