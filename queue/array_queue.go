/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/17 下午6:19
*/
package queue

type Queue struct {
	items []int //数组
	head int
	tail int
}
var myQueue Queue

//入队
func enqueue(item int) bool {
	myQueue.items = append(myQueue.items, item)
	myQueue.tail++
	return true
}

//出队
func dequeue() interface{} {
	//head==tail代表队列为空
	if  myQueue.head == myQueue.tail {
		return nil
	}
	item := myQueue.items[myQueue.head]
	myQueue.head++
	return item
}