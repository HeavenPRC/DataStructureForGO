/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/17 下午2:44
*/
package linkedList

import (
	"testing"
)

func TestLink(t *testing.T) {
	List := NewLinkedList()
	List.InsertToTail(1)
	List.InsertToTail(2)
	List.InsertToTail(3)
	List.InsertToTail(4)
	List.InsertToTail(5)
	List.Print()
	List.InsertToHead(111)
	List.InsertToHead(222)
	List.Print()
}