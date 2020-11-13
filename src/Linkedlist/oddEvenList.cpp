struct ListNode {
      int val;
     ListNode *next;
      ListNode() : val(0), next(nullptr) {}
      ListNode(int x) : val(x), next(nullptr) {}
      ListNode(int x, ListNode *next) : val(x), next(next) {}
  };
 
class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        if(head == nullptr){
          return head;
        }
        ListNode* evenHead = head->next;
        //扫描奇节点
        ListNode* odd = head;
        //扫描偶节点
        ListNode* even = evenHead;
        while (even != nullptr || even->next != nullptr)
        {
          even->next = odd->next;
          even = even->next;
          odd->next = even->next;
          odd = odd->next;
        }
        odd->next = evenHead;
        return head;
    }
};