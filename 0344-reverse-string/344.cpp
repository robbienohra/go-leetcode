#include <vector>

using namespace std;

class Solution {
public:
  void reverseString(vector<char> &s) {

    int n = s.size();

    int L = 0, R = n - 1;

    while (L < R) {

      int temp = s[L];

      s[L] = s[R];
      L++;
      s[R] = temp;
      R--;
    }
  }
};