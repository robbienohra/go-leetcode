#include <vector>

using namespace std;

class Solution {
public:
  vector<int> productExceptSelf(vector<int> &nums) {

    int n = nums.size();

    vector<int> output(n);
    vector<int> L(n);
    vector<int> R(n);

    L[0] = 1;
    R[n - 1] = 1;

    for (int i = 1; i < n; i++) {
      L[i] = L[i - 1] * nums[i - 1];
    }

    for (int i = n - 1; i > 0; i--) {
      R[i - 1] = R[i] * nums[i];
    }

    for (int i = 0; i < n; i++) {
      output[i] = L[i] * R[i];
    }

    return output;
  }
};