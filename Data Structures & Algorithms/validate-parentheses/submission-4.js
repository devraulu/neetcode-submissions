class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(str) {
        const stack = [];

        const bracketsMap = new Map();
        bracketsMap.set(")", "(");
        bracketsMap.set("}", "{");
        bracketsMap.set("]", "[");

        const isOpeningBracket = (char) => "({[".includes(char);

        for (const char of str) {
            if (bracketsMap.has(char)) {
                if (stack.pop() !== bracketsMap.get(char)) return false;
            }
            else if (isOpeningBracket(char)) {
                stack.push(char);
            }
        }

        return stack.length === 0;
    }
}


