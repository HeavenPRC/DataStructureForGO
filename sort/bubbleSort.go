/**
* 冒泡排序
* 每次对一个数据进行冒泡（进行数据交换）如果有一次没有发生冒泡 就可以直接退出算法
* @Author: zhangjian@mioji.com
* @Date: 2019/5/27 上午10:33
*/
package sort

import "fmt"

func BubbleSort(a []int) {
	n := len(a)

	for i := 0;i < n; i++ {
		//标记数据是否发生转移
		flag := false

		for j := 0; j < n - i - 1; j++  {

			if a[j] > a[j + 1] {
				flag = true
				a[j], a[j+1] = a[j+1], a[j]
			}
		}

		if !flag {
			break
		}
	}

	fmt.Println(a)
}
