package main

//利用哈希表O(1)查找时间复杂度的特性
func twoSum(nums []int,target int) []int{
	m := make(map[int]int)
	for k,v := range nums{
		if p,ok := m[target-k];ok{
			return []int{p,k}
		}
		m[v] = k
	}
	return nil
}