### 二分查找的前提
+ 目标函数单调性（单调递增或者单调递减）
+ 存在上下界
+ 能够通过索引访问

#### 代码模板
~~~ go
func binarySearch(arr[]int,target int) int{
	if(arr == nil || len(arr) == 0){
		return -1
	}
	left := 0
	right := len(arr) - 1
	for left<=right{
		mid := left + (right - left)/2
		//fmt.Println(mid)
		if arr[mid] > target{
			right = mid-1
		}else if arr[mid] < target{
			left = mid +1
		}else if arr[mid] == target{
			return mid
		}
	}
	return -1
}
~~~