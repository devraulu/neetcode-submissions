class MinStack {
    stack = []
    minStack = []
    constructor() {

    }

    /**
     * @param {number} val
     * @return {void}
     */
    push(val) {
        this.stack.push(val)
        const min = this.minStack[this.minStack.length - 1];
        if (min >= val || min == undefined || min == null) this.minStack.push(val);
    }

    /**
     * @return {void}
     */
    pop() {
        const popped = this.stack.pop();
        if (popped == this.minStack[this.minStack.length - 1]) this.minStack.pop()
    }

    /**
     * @return {number}
     */
    top() {
        return this.stack[this.stack.length - 1];
    }

    /**
     * @return {number}
     */
    getMin() {
        return this.minStack[this.minStack.length - 1];
    }
}
