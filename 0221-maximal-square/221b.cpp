#include <vector>

using namespace std;

class Solution {
public:
  int maximalSquare(vector<vector<char>> &matrix) {

    int rows = matrix.size();

    if (rows == 0 || matrix[0].empty()) {
      return 0;
    }

    int cols = matrix[0].size();

    vector<int> dp(cols + 1, 0);

    int maxsqlen = 0, prev = 0;

    for (int i = 1; i <= rows; i++) {
      for (int j = 1; j <= cols; j++) {
        int temp = dp[j];
        if (matrix[i - 1][j - 1] == '1') {
          dp[j] = min(min(dp[j - 1], prev), dp[j]) + 1;
          maxsqlen = max(maxsqlen, dp[j]);
        } else {
          dp[j] = 0;
        }
        prev = temp;
      }
    }
    return maxsqlen * maxsqlen;
  }
};