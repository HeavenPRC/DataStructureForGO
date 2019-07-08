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