package main

type Stack struct {
	data *Node
	next *Stack
}

//generateStack 生成空栈
func generateStack() *Stack {
	return &Stack{
		data: nil,
		next: nil,
	}
}

//pushStack 入栈
func pushStack(stack *Stack, data *Node) {
	tmp := new(Stack)
	tmp.data = data
	tmp.next = stack.next
	stack.next = tmp
}

//popStack 出栈
func popStack(stack *Stack) *Node {
	tmp := new(Stack)
	tmp = stack.next
	stack.next = tmp.next
	return tmp.data
}

//emptyStack 判空
func emptyStack(stack *Stack) bool {
	if stack.next == nil {
		return true
	}
	return false
}
