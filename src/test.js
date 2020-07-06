const maxArea = function(nums){
  let i = 0
  let j = nums.length
  let res = 0
  while(i<j){
    res = nums[i] < nums[j] ? Math.max(res,(j - i) * nums[i++]):Math.max(res,(j-i)*nums[j--])
  }
  return res
}