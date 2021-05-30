#include <string>

using namespace std;

class Solution {
public:
  int expand(string &s, int i, int j) {
    int L = i, R = j;
    int n = s.size();

    while (L >= 0 && R < n && s[L] == s[R]) {
      L--;
      R++;
    }

    return R - L - 1;
  }
  string longestPalindrome(string s) {

    int n = s.size();
    int start = 0, end = 0;
    for (int i = 0; i < n; i++) {
      int even = expand(s, i, i);
      int odd = expand(s, i, i + 1);
      int len = max(even, odd);
      int maxLen = end - start;
      if (len > maxLen) {
        start = i - (len - 1) / 2;
        end = i + len / 2;
      }
    }
    return s.substr(start, end + 1 - start);
  }
};