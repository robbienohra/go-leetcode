#include <climits>
#include <string>
#include <unordered_map>

using namespace std;

class Solution {
public:
  int lengthOfLongestSubstringTwoDistinct(string s) {

    int n = s.size();

    if (n < 3) {
      return n;
    }

    int L = 0, R = 0;

    int maxLen = 2;

    unordered_map<char, int> m;

    while (R < n) {
      m[s[R]] = R;

      R++;
      if (m.size() == 3) {

        int leftMost = INT_MAX;

        unordered_map<char, int>::iterator it;

        for (it = m.begin(); it != m.end(); it++) {
          if (it->second < leftMost) {
            leftMost = it->second;
          }
        }

        m.erase(s[leftMost]);
        L = leftMost + 1;
      }

      if (R - L > maxLen) {
        maxLen = R - L;
      }
    }

    return maxLen;
  }
};