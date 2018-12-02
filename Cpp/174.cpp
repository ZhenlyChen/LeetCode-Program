#include<cstdio>
#include<iostream>
#include<vector>

using namespace std;

class Solution {
public:
    int calculateMinimumHP(vector<vector<int>>& dungeon) {
        int row = dungeon.size();
        int col = dungeon[0].size();
        //int hp[row + 1][col + 1];
        vector<vector<int> > hp(row + 1, vector<int>(col + 1, INT_MAX));
        // int minHp[row][col];
        // if (dungeon[0][0] < 0) {
        //     hp[0][0] = 1;
        //     minHp[0][0] = 1 - dungeon[0][0];
        // } else {
        //     minHp[0][0] = 1;
        //     hp[0][0] = dungeon[0][0] + 1;
        // }
        hp[row - 1][col] = 1;
        hp[row][col - 1] = 1;

        for (int i = row - 1; i >= 0; i--) {
          for (int j = col - 1; j >= 0; j--) {
              int need = min(hp[i + 1][j], hp[i][j + 1]) - dungeon[i][j];
              hp[i][j] = need <= 0 ? 1 : need;
          }
        }



        // for (int i = 0; i < row; i++) {
        //     for (int j = 0; j < col; j++) {
        //         if (i == 0 && j != 0) {
        //             int nowHp = hp[i][j - 1] + dungeon[i][j];
        //             if (nowHp >= 1) {
        //                 hp[i][j] = nowHp;
        //                 minHp[i][j] = minHp[i][j - 1];
        //             } else {
        //                 hp[i][j] = 1;
        //                 minHp[i][j] = minHp[i][j - 1] + (1 - nowHp);
        //             }
        //         } else if (j == 0 && i != 0) {
        //             int nowHp = hp[i - 1][j] + dungeon[i][j];
        //             if (nowHp >= 1) {
        //                 hp[i][j] = nowHp;
        //                 minHp[i][j] = minHp[i - 1][j];
        //             } else {
        //                 hp[i][j] = 1;
        //                 minHp[i][j] = minHp[i - 1][j] + (1 - nowHp);
        //             }
        //         } else if (i != 0 && j != 0) {
        //             int nowHp;
        //             int minLast;
        //             if ((minHp[i - 1][j] == minHp[i][j - 1] && hp[i - 1][j] < hp[i][j - 1]) || (minHp[i - 1][j] > minHp[i][j - 1])) {
        //                 minLast = minHp[i][j - 1];
        //                 nowHp = hp[i][j - 1];
        //             } else {
        //                 minLast = minHp[i - 1][j];
        //                 nowHp = hp[i - 1][j];
        //             }
        //             nowHp += dungeon[i][j];
        //             if (nowHp >= 1) {
        //                 hp[i][j] = nowHp;
        //                 minHp[i][j] = minLast;
        //             } else {
        //                 hp[i][j] = 1;
        //                 minHp[i][j] = minLast + (1 - nowHp);
        //             }
        //         }
        //     }
        // }
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                printf("%d ", dungeon[i][j]);
            }
            printf("\n");
        }
        printf("\n");

        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                printf("%d ", hp[i][j]);
            }
            printf("\n");
        }
        printf("\n");
        // for (int i = 0; i < row; i++) {
        //     for (int j = 0; j < col; j++) {
        //         printf("%d ", minHp[i][j]);
        //     }
        //     printf("\n");
        // }

        return hp[0][0];
    }
};

class Solution {
public:
    int maxArea(vector<int>& height) {
        int left = 0;
        int right = height.size() - 1;
        int water = 0;
        int h = 0;
        while (left < right) {
            h = min(height[left], height[right]);
            water = max(water,  h * (right - left));
            while (height[left] <= h && left < right) left++;
            while (height[right] <= h && left < right) right--;
        }
        return water;
    }
};

int main() {
  Solution s;
  cout << "Test Begin" << endl;
  cout << "Input:[[1,-3,3],[0,-2,0],[-3,-3,-3]] \nOutput: 3" << endl;
  cout << "Test Finish" << endl;
}
