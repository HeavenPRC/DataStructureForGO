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

func TestBsearchFirst(t *testing.T) {
	a := []int{1,2,3,4,4,4,4,6,7,7,7,5,6,7,8,9}
	fmt.Println(BsearchFirst(a, 4))
}

func TestBsearchEnd(t *testing.T) {
	a := []int{1,2,3,4,4,4,4,6,7,7,5,6,7,8,9}
	fmt.Println(BsearchEnd(a, 4))
}

func TestBsearchEqFirst(t *testing.T) {
	a := []int{1,2,3,4,4,4,4,6,6,6,6,667,7,5,6,7,8,9}
	fmt.Println(BsearchEnd(a, 4))
}

func TestSearchAb(t *testing.T) {
	fmt.Println("search 开始")
	//a := []int{6,6,7,5,6,7,8,9,1,2,3,4,4,4,4,6,6}
	//fmt.Println("searchAb", SearchAb(a))
	b := []int{6,1,2,3,4,5,56}
	fmt.Println("searchAb", SearchAb(b))
}

func TestMap(t *testing.T) {
	data := map[string]interface{}{
		"k":"val",
	}
	for t := range data {
		fmt.Println(t)
	}
}