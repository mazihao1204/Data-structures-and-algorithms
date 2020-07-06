#### 反转链表

**双指针法**

+ curr 为当前指针 ，pre 为后继指针

~~~go
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode = nil
    curr := head
    for curr != nil{
        // 记录下一个节点
        temp := curr.Next
        //当前指针指向后继指针
        curr.Next = pre
        //后继指针向前移一位
        pre = curr
        //当前指针也向前移一位
        curr = temp
    }
    return pre
}
~~~



**递归**

~~~javascript
const reverseList = function(head){
    if(head === null || head.next === null){
        return head
    }
    let curr = reserseList(head.next)
    head.next.next = head
    head.next = null
    return curr
}
~~~

