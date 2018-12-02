#include <cstring>
#include <iostream>
#include <queue>
#include <utility>
#include <vector>
using namespace std;

class Solution {
 public:
  vector<pair<int, int>> getSkyline(vector<vector<int>>& buildings) {
    vector<pair<int, int>> res;
    int x, height, current = 0, len = buildings.size();
    priority_queue<pair<int, int>>
        heightAndRight;  // 存储高度和右边界（按高度从小到大排序）
    while (current < len || !heightAndRight.empty()) {
      cout << "current: " << current << endl;
      // 当前x轴扫描到的位置
      x = heightAndRight.empty() ? buildings[current][0]
                                 : heightAndRight.top().second;
      cout << "x: " << x << endl;
      if (current >= len || buildings[current][0] > x) {
        while (!heightAndRight.empty() && heightAndRight.top().second <= x) {
          cout << "pop " << heightAndRight.top().first << endl;
          heightAndRight.pop();  // 弹出已处理完的矩形
        }
      } else {                      // 处理左边界重叠的情况
        x = buildings[current][0];  // 当前左边界
        cout << "x: " << x << endl;
        while (current < len &&
               buildings[current][0] == x) {  // 处理左边界重叠的情况
          heightAndRight.push(
              make_pair(buildings[current][2], buildings[current][1]));
          cout << "add " << buildings[current][2] << endl;
          current++;
        }
      }
      int height = heightAndRight.empty() ? 0 : heightAndRight.top().first;
      cout << "height: " << height << endl;
      if (res.empty() || res.back().second != height) {
        res.push_back(make_pair(x, height));
        cout << "push " << x << " " << height << endl;
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  cout << "Test Begin" << endl;
  cout << "Input:[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]] " << endl;
  cout << "Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]" << endl;
  vector<vector<int>> input{vector<int>{2, 9, 10}, vector<int>{3, 7, 15},
                            vector<int>{5, 12, 12}, vector<int>{15, 20, 10},
                            vector<int>{19, 24, 8}};
  vector<pair<int, int>> res = s.getSkyline(input);
  for (auto r : res) {
    cout << r.first << " " << r.second << endl;
  }
  cout << "Test Finish" << endl;
}
