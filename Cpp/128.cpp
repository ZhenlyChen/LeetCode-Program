class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        unordered_set<int> numSet;
        for(int num : nums) {
            numSet.insert(num);
        }
        int maxSeqLen = 0;
        for(int num : numSet) {
            if (numSet.find(num - 1) == numSet.end()) {
                int currentLen = 1;
                int currentNum = num;
                while (numSet.find(currentNum + 1) != numSet.end()) {
                    currentLen++;
                    currentNum++;
                }
                if (currentLen > maxSeqLen) {
                    maxSeqLen = currentLen;
                }
            }
        }
        return maxSeqLen;
    }
};
