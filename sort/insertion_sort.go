/**
* 插入排序
* @Author: zhangjian@mioji.com
* @Date: 2019/5/24 下午5:25
*/
package sort

import "fmt"

func InsertionSort(a []int) {
	n := len(a)
	//打印退出
	if n <= 1 {
		fmt.Println(a)
		return
	}

	for i:= 1; i < n; i++ {
		value := a[i]
		j := i - 1
		//查找插入的位置 一个循环处理一个数据
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}

	fmt.Println(a)
}

