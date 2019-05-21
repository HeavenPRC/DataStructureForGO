/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/20 下午5:35
*/
package queue

import (
	"fmt"
	"testing"
)
var MyListQueue LinkQueue

//初始化队列
func init() {
	MyListQueue.head = &ListNode{nil, 0}
	MyListQueue.tail = MyListQueue.head
}

//打印链表
func (this *LinkQueue) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func TestLink(t *testing.T) {
	MyListQueue.lenqueue(1)
	MyListQueue.lenqueue(2)
	MyListQueue.lenqueue(3)
	MyListQueue.lenqueue(4)
	MyListQueue.Print()
	fmt.Println(MyListQueue.ldequeue())
	MyListQueue.Print()
	fmt.Println(MyListQueue.ldequeue())
	MyListQueue.Print()
	fmt.Println(MyListQueue.ldequeue())
	MyListQueue.Print()
	fmt.Println(MyListQueue.ldequeue())
	MyListQueue.Print()
	fmt.Println(MyListQueue.ldequeue())
	MyListQueue.Print()
	fmt.Println(MyListQueue.ldequeue())
	MyListQueue.Print()

}