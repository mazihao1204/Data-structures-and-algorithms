#include<vector>
using namespace std;

class solution{
  public:
    //移动0
    void moveZeros(vector<int>&nums){
      int left = 0,right = 0,n = nums.size();
      while(right < n){
        if(nums[right] != 0){
          swap(nums[left],nums[right]);
          right++;
        }
        left++;
      }
    }


    //移除元素,
    int moveElement(vector<int>& nums,int val){
        int i = 0,n = nums.size();
        while(i < n){
          if(nums[i] == val){
            nums[i] = nums[n-1]; 
          }else{
            i++;
          }
        }
        return n;
    }
};