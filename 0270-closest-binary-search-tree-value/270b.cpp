#include <climits>
#include <math.h>
#include <stack>
#include <vector>

using namespace std;

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

class Solution {
public:
  int closestValue(TreeNode *root, double target) {

    int min = INT_MIN;

    stack<TreeNode *> travStack;

    TreeNode *curr = root;

    while (curr != nullptr || !travStack.empty()) {
      while (curr != nullptr) {
        travStack.push(curr);
        curr = curr->left;
      }

      curr = travStack.top();
      travStack.pop();

      if (abs(static_cast<float>(curr->val) - target) <
          abs(static_cast<float>(min) - target)) {
        min = curr->val;
      }
      curr = curr->right;
    }

    return min;
  }
};