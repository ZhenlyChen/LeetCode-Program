#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int numDistinct(string s, string t) {
    vector<vector<int>> dis(s.length() + 1, vector<int> (t.length() + 1, 0));
    for (int i = 0; i <= t.length(); i++) {
      dis[0][i] = 1;
    }
    for (int i = 0; i <= s.length(); i++) {
      for (int j = 0; j <= t.length(); j++) {
        if (s[i] == t[j]) {
          dis[i + 1][j + 1] = dis[i][j] + dis[i + 1][j];
        } else {
          dis[i + 1][j + 1] = dis[i + 1][j];
        }
      }
    }
    return dis[s.length()][t.length()];
  }
};

int main() {
  Solution s;
  cout << "Test Begin" << endl;
  cout << "Input: S = \"rabbbit\", T = \"rabbit\"\nOutput: 3" << endl;
  cout << s.numDistinct("rabbbit", "rabbit") << endl;
  cout << "Input: S = \"babgbag\", T = \"bag\"\nOutput: 5" << endl;
  cout << s.numDistinct("babgbag", "bag") << endl;
  cout << "Test Finish" << endl;
}
