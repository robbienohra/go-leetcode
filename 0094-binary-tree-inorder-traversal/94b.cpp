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
    vector<int> res;

    stack<TreeNode *> stack;

    TreeNode *curr = root;

    while (curr != nullptr || !stack.empty()) {

      while (curr != nullptr) {
        stack.push(curr);
        curr = curr->left;
      }
      // leftmost child would have been encountered at this point

      curr = stack.top();

      stack.pop();

      res.push_back(curr->val);

      // if this is null (i.e. the leftmost child has no right child)
      // then at the next iteration the initial loop above will be skipped
      // and we simply proceed to access the parent of this current node
      // which is the next leftmost child
      // if this is not null then at the next iteration we proceed to
      // find the leftmost child of the right subtree

      curr = curr->right;
    }

    return res;
  }
};