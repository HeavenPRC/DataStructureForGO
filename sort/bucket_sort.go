/**
* 桶排序 时间复杂度 在理想情况下是On 在最坏情况下会退化到nlogn 空间复杂度On 由数据规模决定
* 1.对数据进行区段分桶 ，桶是天然有序的
* 2.对每个桶内的数据进行快排
* 3.将数据依次取出覆盖拷贝到初始数据中
* @Author: zhangjian@mioji.com
* @Date: 2019/6/12 下午3:26
*/
package sort

import "fmt"

// 获取待排序数组中的最大值
func getMax(a []int)int{
	max := a[0]
	for i:=1; i<len(a); i++{
		if a[i] > max{
			max = a[i]
		}
	}
	return max
}


func BucketSort(a []int)  {
	num := len(a)
	if num <= 1{
		return
	}
	max := getMax(a)
	buckets := make([][]int, num)  // 二维切片

	//先把
	index :=0
	for i:=0; i< num; i++{
		index = a[i]*(num -1) / max  // 桶序号
		buckets[index] = append(buckets[index],a[i]) // 加入对应的桶中
	}

	tmpPos := 0  // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			QuickSort(buckets[i], 0 ,len(buckets[i]))  // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}

}


// 桶排序简单实现
func BucketSortSimple(source []int)  {
	if len(source)<=1{
		return
	}
	array := make([]int, getMax(source)+1)
	for i:=0; i<len(source); i++{
		array[source[i]] ++
	}
	fmt.Println(array)
	c := make([]int,0)
	for i:=0; i<len(array); i++{
		for array[i] != 0 {
			c = append(c, i)
			array[i] --
		}
	}
	copy(source,c)

}