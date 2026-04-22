class Solution {
    /**
     * @param {number[]} temperatures
     * @return {number[]}
     */
    dailyTemperatures(temperatures) {
        const stack = [];
        const results = [];

        for (let i = 0; i < temperatures.length; i++) {
            const temp = temperatures[i];
            results[i] = 0;

            while (temp > temperatures[stack[stack.length - 1]]) {
                const j = stack.pop();
                results[j] = i - j;
            }

            stack.push(i);
        }

        return results;
    }
}
