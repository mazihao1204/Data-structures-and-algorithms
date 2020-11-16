#include <vector>
using namespace std;

struct ListNode {
      int val;
     ListNode *next;
      ListNode() : val(0), next(nullptr) {}
      ListNode(int x) : val(x), next(nullptr) {}
      ListNode(int x, ListNode *next) : val(x), next(next) {}
  };
 
class Solution {
public:
    //奇偶链表
    ListNode* oddEvenList(ListNode* head) {
        if(head == nullptr){
          return head;
        }
        ListNode* evenHead = head->next;
        //扫描奇节点
        ListNode* odd = head;
        //扫描偶节点
        ListNode* even = evenHead;
        while (even != nullptr && even->next != nullptr)
        {
          even->next = odd->next;
          even = even->next;
          odd->next = even->next;
          odd = odd->next;
        }
        odd->next = evenHead;
        return head;
    }

    //分割链表
    vector<ListNode*> splitListToParts(ListNode* root, int k) {
      ListNode* p = root;
      int size = 0;
      while(p != nullptr){
        size++;
        p = p->next;
      }
      int averLength = size/k;
      int remainder = size%k;
      vector<ListNode*> res(k,nullptr);
      
    }
};