/**
* 计数排序 桶排序的一种特殊情况 排序的n个数据范围不大
* 1.遍历数据进行计数 桶的键和值相同
* 2.如过不考虑排序稳定性直接遍历辅助变量中的数据即可,否则对辅助数组进行遍历求和
* 3.从右边开始遍历初始数据 ，放入中间数组，最后把中间数组copy到初始数组
* @Author: zhangjian@mioji.com
* @Date: 2019/7/7 下午4:32
*/
package sort

import (
	"fmt"
)

func CountingSort(a []int) {
	max := getMax(a)
	//初始化中间数组和计数数组
	temp_a, count_a := make([]int, len(a)), make([]int, max + 1)
	//只有一个接收者是k
	for _,v := range a {
		fmt.Println("v",v)
		count_a[v]++
	}
	//对计数数组进行迭代求值
	counts := 0
	for k, v := range count_a{
		if v != 0 {
			counts, count_a[k] = v + counts, v + counts
		}
	}
	fmt.Println("计数器", count_a)

	for _, v := range a {
		true_k := count_a[v] - 1
		temp_a[true_k] = v
		count_a[v]--
	}
	//把处理好的数组拷贝回原始数组
	copy(a[:], temp_a[:])
	fmt.Println(a)
}
