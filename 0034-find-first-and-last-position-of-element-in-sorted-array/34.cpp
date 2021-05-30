#include <math.h>
#include <vector>

using namespace std;

class Solution {
public:
  int first(vector<int> &nums, int target) {

    int n = nums.size();

    if (n == 0) {
      return -1;
    }

    int L = 0, R = n - 1;

    while (L != R) {
      int m = floor(L + static_cast<float>(R - L) / 2);

      if (nums[m] >= target) {

        R = m;

      } else {

        L = m + 1;
      }
    }

    if (nums[L] == target) {
      return L;
    }

    return -1;
  }

  int last(vector<int> &nums, int target) {
    int n = nums.size();

    if (n == 0) {
      return -1;
    }

    int L = 0, R = n - 1;

    while (L != R) {
      int m = ceil(L + static_cast<float>(R - L) / 2);
      if (nums[m] <= target) {
        L = m;
      } else {
        R = m - 1;
      }
    }

    if (nums[L] == target) {
      return L;
    }

    return -1;
  }
  vector<int> searchRange(vector<int> &nums, int target) {

    int a = first(nums, target);
    int b = last(nums, target);

    return {a, b};
  }
};