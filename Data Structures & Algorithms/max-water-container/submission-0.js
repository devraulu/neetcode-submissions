class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
        let left = 0, right = heights.length - 1, res = 0;

        while (left < right) {
            const area = Math.min(heights[left], heights[right]) * Math.abs(left - right);
            res = Math.max(area, res);
            if (heights[left] <= heights[right]) {
                left++;
            } else {
                right--;
            }
        }

        return res;
    }
}
