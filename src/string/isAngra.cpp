#include<vector>
#include<string>
using namespace std;

class Solution{
  public:

    //异位词
    bool isAngram(string s,string t){
      if(s.length() != t.length()) return false;
      vector<int> res(26,0);
      for(auto& ch : s){
        res[ch-'a']++;
      }
      for(auto& ch : t){
        res[ch-'a']--;
        if(res[ch-'a'] < 0){
          return false;
        }
      }
      return true;
    }
};