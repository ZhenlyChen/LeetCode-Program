#include <iostream>
#include <string>
using namespace std;

class Solution {
 public:
  int minDistance(string word1, string word2) {
    int len1 = word1.length();
    int len2 = word2.length();
    int mat[len1 + 1][len2 + 1];
    for (int i = 0; i <= len1; i++) {
      mat[i][0] = i;
    }
    for (int j = 0; j <= len2; j++) {
      mat[0][j] = j;
    }
    for (int i = 1; i <= len1; i++) {
      for (int j = 1; j <= len2; j++) {
        if (word1[i - 1] == word2[j - 1]) {
          mat[i][j] = mat[i - 1][j - 1];
        } else {
           mat[i][j] = min(mat[i - 1][j - 1], min(mat[i - 1][j],  mat[i][j - 1])) + 1;
        }
      }
    }
    // for (int i = 0; i <= len1; i++) {
    //   for (int j = 0; j <= len2; j++) {
    //     cout << mat[i][j] << "\t";
    //   }
    //   cout << endl;
    // }
    return mat[len1][len2];
  }
};

int main() {
  Solution s;
  cout << "Test Begin" << endl;
  cout << "word1 = \"horse\", word2 = \"ros\"" << endl;
  cout << s.minDistance("horse", "ros") << endl;
  cout << "word1 = \"intention\", word2 = \"execution\"" << endl;
  cout << s.minDistance("intention", "execution") << endl;
  cout << "word1 = \"exponential\", word2 = \"polynomial\"" << endl;
  cout << s.minDistance("exponential", "polynomial") << endl;
  cout << "Test Finish" << endl;
}
