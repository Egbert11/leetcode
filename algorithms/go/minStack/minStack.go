// Source: https://leetcode.com/problems/min-stack/
// Author: Egbert11
// Date: 2019-05-07

/*
	Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

	push(x) -- Push element x onto stack.
	pop() -- Removes the element on top of the stack.
	top() -- Get the top element.
	getMin() -- Retrieve the minimum element in the stack.
	Example:
		MinStack minStack = new MinStack();
		minStack.push(-2);
		minStack.push(0);
		minStack.push(-3);
		minStack.getMin();   --> Returns -3.
		minStack.pop();
		minStack.top();      --> Returns 0.
		minStack.getMin();   --> Returns -2.
*/

package main

import "fmt"

type MinStack struct {
	stack []int
	minSlice []int
	index int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	var minStack = new(MinStack)
	minStack.index = -1
	return *minStack
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	var minValue = x
	if len(this.minSlice) > 0 {
		if x > this.minSlice[this.index] {
			minValue = this.minSlice[this.index]
		}
	}
	this.minSlice = append(this.minSlice, minValue)
	this.index += 1
}


func (this *MinStack) Pop()  {
	if this.index < 0 {
		return
	}
	this.stack = this.stack[:this.index]
	this.minSlice = this.minSlice[:this.index]
	this.index -= 1
}


func (this *MinStack) Top() int {
	return this.stack[this.index]
}


func (this *MinStack) GetMin() int {
	return this.minSlice[this.index]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main()  {
	obj := Constructor();
	obj.Push(-2);
	obj.Push(0);
	obj.Push(-3);
	fmt.Println(obj.GetMin());
	obj.Pop();
	fmt.Println(obj.Top());
	fmt.Println(obj.GetMin());
}
