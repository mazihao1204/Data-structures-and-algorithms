/**
 * 实现 int sqrt(int x) 函数。
   计算并返回 x 的平方根，其中 x 是非负整数。
   由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 */

 //牛顿迭代法
function sqrtsOne(x){
  let r = x
  while(r*r > x){
    r = ((r + x/r) / 2) | 0
  }
  return r
}


//二分查找
var mySqrt = function (x) {
  if (x < 2) return x
  let left = 1, right = x >> 1
  while (left + 1 < right) {
    let mid = left + ((right - left) >> 1)
    if (mid === x / mid) {
      return mid
    } else if (mid < x / mid) {
      left = mid
    } else {
      right = mid
    }
  }
  return right > x / right ? left : right
};
