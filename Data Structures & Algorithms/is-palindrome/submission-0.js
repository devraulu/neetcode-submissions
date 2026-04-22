class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        let left =  0, right = s.length - 1;
        const regex = /\s|\W/;
        while(left < right) {
            if(regex.test(s[left])) {
                left++;
                continue;
            }
            if(regex.test(s[right])){
                right--;
               continue; 
            }
            if(s[left].toLowerCase() != s[right].toLowerCase()) return false;
            left++;
            right--;
        }
        return true;
    }
}
