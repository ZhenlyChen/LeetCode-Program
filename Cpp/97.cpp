#include <iostream>
#include <string>
#include <cstring>
using namespace std;

class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
      if (s3.length() != s1.length() + s2.length()) return false;
      bool mat[s1.length() + 1][s2.length() + 1];
      memset(mat, 0, sizeof(mat));
      for (int i = 0; i <= s1.length(); i++) {
        for (int j = 0; j <= s2.length(); j++) {
          if (i == 0 && j == 0) {
            mat[i][j] = true;
          } else if (i == 0) {
            mat[i][j] = mat[i][j - 1] && s2[j - 1] == s3[i + j - 1];
          } else if (j == 0) {
            mat[i][j] = mat[i - 1][j] && s1[i - 1] == s3[i + j - 1];
          } else {
            mat[i][j] = (mat[i][j - 1] && s2[j - 1] == s3[i + j - 1]) || (mat[i - 1][j] && s1[i - 1] == s3[i + j - 1]);
          }
          cout << mat[i][j] << '\t';
        }
        cout << endl;
      }
      return mat[s1.length()][s2.length()];
    }
};

int main() {
  Solution s;
  cout << "Test Begin" << endl;
  cout << "Input: s1 = \"aabcc\", s2 = \"dbbca\", s3 = \"aadbbcbcac\" \nOutput: true" << endl;
  cout << s.isInterleave("aabcc", "dbbca", "aadbbcbcac") << endl;
  cout << "Input: s1 = \"aabcc\", s2 = \"dbbca\", s3 = \"aadbbbaccc\" \nOutput: false" << endl;
  cout << s.isInterleave("aabcc", "dbbca", "aadbbbaccc") << endl;
  cout << "Test Finish" << endl;
}
