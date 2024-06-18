package flex

type Stack[T any] struct {
	top      int64
	length   uint32
	elements []T
}

func NewStack[T any](length uint32) *Stack[T] {
	return &Stack[T]{
		top:      -1,
		length:   length,
		elements: make([]T, length),
	}
}

func (stack *Stack[T]) Push(item T) bool {
	if stack.Full() {
		return false
	}
	stack.top++
	stack.elements[stack.top] = item
	return true
}

func (stack *Stack[T]) Peek() (item T, ok bool) {
	if stack.Empty() {
		return
	}
	return stack.elements[stack.top], true
}

func (stack *Stack[T]) Pop() (item T, ok bool) {
	if stack.Empty() {
		return
	}
	item = stack.elements[stack.top]
	stack.top--
	return item, true
}

func (stack *Stack[T]) Empty() bool {
	return stack.top == -1
}

func (stack *Stack[T]) Full() bool {
	return stack.top+1 == int64(stack.length)
}

func (stack *Stack[T]) Length() int64 {
	return stack.top + 1
}
