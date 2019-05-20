/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/20 下午5:35
*/
package queue

import (
	"fmt"
	"testing"
)

func init() {
	myQueue = Queue{
		items: []int{},
		head: 0,
		tail: 0,
	}
}

func TestLink(t *testing.T) {
	enqueue(1)
	enqueue(2)
	enqueue(3)
	enqueue(4)
	fmt.Println(dequeue())
	fmt.Println(dequeue())
	fmt.Println(dequeue())
	fmt.Println(dequeue())
	fmt.Println(dequeue())
}