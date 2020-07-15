package main

import (
	"sort"
	"strconv"
	"strings"
)
/**
 	leetcode179  最大数
	给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
	剑指offer 将数组排成最小的数,两题都是同一思路
	思路：对于x+y > y+x 则有 x > y
		反之：x+y < y+x 则有x < y
 */
func largestNumber(nums []int) string{
	ss := make([]string,len(nums))
	for i,num := range nums{
		//整数转为字符串 strconv.Itoa()
		ss[i] = strconv.Itoa(num)
	}
	//自定义排序规则
	sort.Slice(ss,func(i,j int) bool{
		return ss[i]+ss[j] <= ss[j] + ss[i]
	})
	res := strings.Join(ss,"")
	//如果第一位为0，那么后面都是0，直接返回0 注意：字符串比较是按字节比较
	if res[0] == '0' {
		return "0"
	}
	return res
}

func minNumber(nums []int) string{
	ss := make([]string,len(nums))
	for i,num := range nums{
		ss[i] = strconv.Itoa(num)
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i] + ss[j] <= ss[j] + ss[i]
	})
	res := strings.Join(ss,"")
	return res
}
