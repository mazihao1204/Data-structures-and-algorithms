/**
 *  leetcode 206 反转链表
 *  两种方法：
 *  1.迭代法
 *  2.递归
 */

 //迭代
 const reverseListOne = function(head){
  let curr = head
  let prev = null
  while(curr){
    let next = curr.next
    curr.next = prev
    prev = curr
    curr = next
  }
  return head
 }

 //递归
 const reverseListTwo = function(head){
  if(head === null || head.next === null) return head
  let p = reverseListTwo(head.next)
  head.next.next = head
  head.next = null
  return p
 }