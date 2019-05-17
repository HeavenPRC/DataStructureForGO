/**
* @Author: zhangjian@mioji.com
* @Date: 2019/5/17 下午2:37
*/
package stack

//没必要用数组
type ArrayStack struct {
	items []interface{}
	count int
}

//入栈
func (a *ArrayStack) push(item interface{}) bool {
	a.count++
	a.items = append(a.items, item)
	return true
}

//出栈
func (a *ArrayStack) pop() interface{} {
	if a.count == 0 {
		return nil
	}
	data := a.items[a.count - 1]
	a.count--
	return data
}
