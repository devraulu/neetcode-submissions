#include <cctype>

class Solution {
public:
    bool isPalindrome(string s) {
        int n = s.length();
        int l = 0, r = n - 1;
        while (l < r) {
            if(!std::isalnum(static_cast<unsigned char>(s[r]))) {
                r--;
                continue;
            } 
           if (!std::isalnum(static_cast<unsigned char>(s[l]))){
                l++;
                continue;
            }
            if (std::tolower(s[l]) != std::tolower(s[r])) {
                return false;
            }
            r--;
            l++;
        }
        return true;
    }
};
