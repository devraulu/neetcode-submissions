class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
       const frequencyMap = new Map();
       for(const n of nums) {
            frequencyMap.set(n, (frequencyMap.get(n) || 0) + 1);
       }

        const bucket = Array.from({ length: nums.length + 1}, () => []);

        for(const [n, count] of frequencyMap){
            bucket[count].push(n)
        }
        
        const topK = [];

        for(let i = bucket.length - 1; i >= 0; i--) {
            for(const num of bucket[i]){
                console.log(i, topK);
                topK.push(num);
                if(topK.length === k) 
                    return topK;
            }
        }

        return topK;
    }
}
