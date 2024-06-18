package flex

type Stack[T any] struct {
	top      int64
	length   uint32
	elements []T
}

// NewStack initializes a new stack with the specified maximum length.
func NewStack[T any](length uint32) *Stack[T] {
	return &Stack[T]{
		top:      -1,
		length:   length,
		elements: make([]T, length), // Initialize the elements slice with zero values
	}
}

// Push adds a new item onto the stack.
func (stack *Stack[T]) Push(item T) bool {
	if stack.Full() {
		return false
	}
	stack.top++
	stack.elements[stack.top] = item
	return true
}

// DynamicPush adds a new item onto the stack and dynamically expands the stack if it's full.
func (stack *Stack[T]) DynamicPush(item T) bool {
	if stack.Full() {
		var (
			oldLength   = stack.length
			newLength   = oldLength * 2
			newElements = make([]T, newLength)
		)
		// Copy existing elements to the new slice
		for i := uint32(0); i < oldLength; i++ {
			newElements[i] = stack.elements[i]
		}
		stack.elements = newElements
		stack.length = newLength
	}
	stack.top++
	stack.elements[stack.top] = item
	return true
}

// Peek returns the top item of the stack without removing it.
func (stack *Stack[T]) Peek() (item T, ok bool) {
	if stack.Empty() {
		return
	}
	return stack.elements[stack.top], true
}

// Pop removes and returns the top item from the stack.
func (stack *Stack[T]) Pop() (item T, ok bool) {
	if stack.Empty() {
		return
	}
	item = stack.elements[stack.top]
	stack.top--
	return item, true
}

// Empty checks if the stack is empty.
func (stack *Stack[T]) Empty() bool {
	return stack.top == -1
}

// Full checks if the stack is full.
func (stack *Stack[T]) Full() bool {
	return stack.top+1 == int64(stack.length)
}

// Length returns the current number of items in the stack.
func (stack *Stack[T]) Length() int64 {
	return stack.top + 1
}
