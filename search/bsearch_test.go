/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/8 下午12:23
*/
package search

import (
	"fmt"
	"testing"
)

func TestBsearch(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(Bsearch(a, 1))
}

func TestBsearch2(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(Bsearch2(a, 3))
}