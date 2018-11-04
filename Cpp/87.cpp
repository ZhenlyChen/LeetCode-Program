#include <string>
#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    bool isScramble(string s1, string s2) {
      // vector<int> part;

    }
};

int main() {
  Solution s;
  cout << "Test Begin" << endl;
  cout << "s1 = \"great\", s2 = \"rgeat\"" << endl;
  cout << s.isScramble("great", "rgeat") << endl;
  cout << "s1 = \"abcde\", s2 = \"caebd\"" << endl;
  cout << s.isScramble("abcde", "caebd") << endl;
  cout << "s1 = \"rgtae\", s2 = \"great\"" << endl;
  cout << s.isScramble("rgtae", "great") << endl;
  cout << "Test Finish" << endl;
}
