#include<vector>
using namespace std;

struct ListNode
{
  int val;
  ListNode* next;
  ListNode(int x): val(x),next(NULL){}
};

class Solution{
  public:

  //反转链表
    ListNode* reverseList(ListNode* head){
      ListNode* curr = head,*prev = NULL;
      while(curr != NULL){
        ListNode* next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
      }
      return prev->next; 
    }

    //判断链表是否有环
    bool hasCycle(ListNode* head){
      if(head == NULL || head->next == NULL) return false;
      ListNode* slow = head,*fast = head;
      while(fast->next == NULL|| fast->next->next == NULL){
        if(slow == fast){
          return true;
        }
        slow = slow->next;
        fast = fast->next->next;
      }
      return false;
    }

    bool hasCycle1(ListNode* head){
      if(head == nullptr || head->next == nullptr){
        return false;
      }
      ListNode* slow = head,*fast = head->next;
      while(slow != fast){
        if(fast == nullptr || fast->next == nullptr)return false;
        slow = slow->next;
        fast = fast->next->next;
      }
      return true;
    }
};
