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
};