#include <vector>

using namespace std;

class Solution {
public:
  int maximalSquare(vector<vector<char>> &matrix) {

    int H = matrix.size();

    if (H == 0 || matrix[0].empty()) {
      return 0;
    }

    int W = matrix[0].size();

    vector<vector<int>> dp(H, vector<int>(W));
  }
};