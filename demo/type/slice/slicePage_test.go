package main

import (
	"fmt"
	"testing"
)

func Test_paging(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	arr, err := cursor(2, 5, data)
	fmt.Println("arr:", arr, "err:", err)

	fmt.Println(paging(2, 5, data))
}

func paging(pn int, ps int, data []int) (res []int) {
	var start, end int
	if pn > 1 {
		start = (pn - 1) * ps
	}
	end = pn * ps
	total := len(data)
	if total == 0 {
		return
	}
	switch {
	case total <= start:
		res = make([]int, 0)
	case total <= end:
		res = data[start:total]
	default:
		res = data[start:end]
	}
	return
}

func cursor(start int, ps int, data []int) ([]int, string) {
	var arr []int
	end := ps + start - 1
	switch {
	case start > len(data):
		return arr, "初始元素位置大于查询数组长度"
	case start+ps-1 > len(data):
		arr = data[start-1 : len(data)]
		return arr, "已经到了最后一页"
	//	case ps > len(data):
	//	arr = data
	case start < len(data) && end <= len(data) && ps <= len(data):
		println(start, end)
		arr = data[start-1 : end]
		return arr, ""
	}
	return arr, "未知情况"
}
