### 二分查找的前提
+ 目标函数单调性（单调递增或者单调递减）
+ 存在上下界
+ 能够通过索引访问

#### 代码模板
~~~ js
let left = 0
let right = 0
let len = arr.length
while(left <= right){
  mid = (right + left) / 2
  if(arr[mid] === target){
    // find the target
    // break or return result
  }
  else if(mid < target){
    left = mid + 1
  }
  else{
    right = mid - 1
  }
}
~~~