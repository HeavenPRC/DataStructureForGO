/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/17 下午6:19
*/
package queue

import "fmt"

type Queue struct {
	items []int //数组
	head int
	tail int
}
var myQueue Queue

func main() {
	fmt.Println(myQueue)
	xx := make([]int, 8)
	xx[0] =1;
	xx[0] =1;xx[0] =1;
	xx[0] =1;xx[0] =1;xx[0] =1;xx[0] =1;
	xx[0] =1;
	xx[0] =1;xx[9] =1;
	fmt.Println(xx)
}

func init() {
	myQueue = Queue{
		items: []int{},
		head: 0,
		tail: 0,
	}
}

//入队
func enqueue(item int) bool {
	myQueue.items[myQueue.tail] = item
	myQueue.tail++
	return true
}

//出队
func dequeue() int {
	//head==tail代表队列为空
	if head = tail {
		
	}
}