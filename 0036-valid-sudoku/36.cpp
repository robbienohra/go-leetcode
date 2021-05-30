#include <set>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
  int block(int row, int col) { return (row / 3) * 3 + col / 3; }

  bool isValidSudoku(vector<vector<char>> &board) {
    unordered_map<int, unordered_map<char, bool>> cols, rows, blocks;

    for (int i = 0; i < 9; i++) {
      for (int j = 0; j < 9; j++) {

        int cell = board[i][j];

        if (cell == '.') {
          continue;
        }

        int k = block(i, j);

        bool row = rows[i][cell], col = cols[j][cell], block = blocks[k][cell];

        if (block || row || col) {
          return false;
        }

        rows[i][cell] = true;
        cols[j][cell] = true;
        blocks[k][cell] = true;
      }
    }

    return true;
  }
};