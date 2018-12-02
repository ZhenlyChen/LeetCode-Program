class Solution {
public:
    int romanToInt(string s) {
        if (s.length() == 0) return 0;
        int last = getNum(s[0]);
        int sum = last;
        for (int i = 1; i < s.length(); i++) {
            int now = getNum(s[i]);
            if (now > last) {
                sum += now - last * 2;
                last = now;
            } else {
                last = now;
                sum += now;
            }
        }
        return sum;
    }

    int getNum(char s) {
        switch(s) {
            case 'I': return 1;
            case 'V': return 5;
            case 'X': return 10;
            case 'L': return 50;
            case 'C': return 100;
            case 'D': return 500;
            case 'M': return 1000;
        }
    }
};
