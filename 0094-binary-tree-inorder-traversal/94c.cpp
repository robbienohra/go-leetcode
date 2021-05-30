struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

#include <stack>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> inorderTraversal(TreeNode *root) {
    TreeNode *p = root;

    vector<int> res;

    stack<TreeNode *> travStack;

    while (p != nullptr) {
      while (p != nullptr) {
        // push nodes onto the stack in reverse access order

        if (p->right != nullptr) {
          travStack.push(p->right);
        }

        travStack.push(p);

        p = p->left;
      }

      // at this point we would have encountered the leftmost left child

      p = travStack.top();
      travStack.pop();

      while (!travStack.empty() && p->right == nullptr) {
        res.push_back(p->val);
        p = travStack.top();
        travStack.pop();
      }

      res.push_back(p->val);

      if (!travStack.empty()) {
        p = travStack.top();
        travStack.pop();
      } else {
        p = nullptr;
      }
    }

    return res;
  }
};