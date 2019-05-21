/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/20 下午6:15
*/
package queue

//定义一个结点
type ListNode struct {
	next *ListNode
	value interface{}
}

type LinkQueue struct {
	head *ListNode
	tail *ListNode 
}

func (MyListQueue *LinkQueue) lenqueue(item interface{})  {
	code :=  &ListNode{nil, item}
	MyListQueue.tail.next = code
	MyListQueue.tail = code

	if MyListQueue.head.next ==  nil {
		MyListQueue.tail.next = MyListQueue.tail
	}

}

func (MyListQueue *LinkQueue) ldequeue() interface{} {
	if MyListQueue.head.next ==  nil {
		return nil
	}
	item := MyListQueue.head.next.value
	MyListQueue.head.next = MyListQueue.head.next.next
	return item
}