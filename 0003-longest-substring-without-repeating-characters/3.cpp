#include <string>
#include <unordered_map>

using namespace std;

class Solution {
public:
  int lengthOfLongestSubstring(string s) {
    int n = s.size();

    int ans = 0;

    unordered_map<char, int> m;

    int i = 0;

    for (int j = 0; j < n; j++) {

      auto pos = m.find(s[j]);

      if (pos != m.end()) {
        if (pos->second + 1 >= i) {
          i = pos->second + 1;
        }
      }

      if (j - i + 1 >= ans) {
        ans = j - i + 1;
      }

      m[s[j]] = j;
    }
    return ans;
  }
};