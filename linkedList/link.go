/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/17 下午2:32
*/
package linkedList

import "fmt"

/**
 * 单链表实现
 */

//定义一个结点
type ListNode struct {
	next *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

//获取下一个节点
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

//获取节点的值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

// 建立哨兵结点
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}
/**
   在参考节点后插入节点
   p 参考节点
   v 新节点的值
 */
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	p.next, newNode.next = newNode, p.next
	this.length++

	return true
}

/**
 *在参考节点前插入节点
 */
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	//从头开始遍历寻在油腻的师姐 上一个节点
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur //pre节点指向当前
		cur = cur.next //cur指向下一个节点
	}
	//遍历到尾节点
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next,newNode.next = newNode, cur
	this.length++
	return true
}

//列表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (this *LinkedList) DeleteTail() bool {
	cur := this.head.next
	var pre, last *ListNode

	if cur == nil {
		return false
	}
	for cur != nil {
		pre,cur = cur, cur.next
	}
	if 	pre == nil {
		return false
	}
	cur = this.head.next

	for cur != nil {
		last, cur = cur, cur.next

		if cur == pre {
			last.next = nil
		}
	}
	return true
}