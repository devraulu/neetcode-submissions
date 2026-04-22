class Solution {
    /**
     * @param {string[]} tokens
     * @return {number}
     */
    evalRPN(tokens) {
        const stack = [];
        const n = tokens.length;

        const operators = new Map([
            ["+", (a, b) => a + b],
            ["-", (a, b) => a - b],
            ["/", (a, b) => Math.trunc(a / b)],
            ["*", (a, b) => a * b]
        ]);

        for (let i = 0; i < n; i++) {
            const operand = tokens[i]
            if (operators.has(operand)) {
                const [b, a] = [stack.pop(), stack.pop()];
                stack.push(operators.get(operand)(a, b))
            } else {
                stack.push(Number(operand))
            }
        }

        return stack.pop();
    }
}
