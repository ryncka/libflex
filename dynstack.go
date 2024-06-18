package flex

type DynamicStack[T any] struct {
	stack *Stack[T]
}

func NewDynamicStack[T any](initialLength uint32) *DynamicStack[T] {
	return &DynamicStack[T]{
		stack: NewStack[T](initialLength),
	}
}

func (dynstack *DynamicStack[T]) Push(item T) bool {
	if dynstack.full() {
		var (
			oldLength   = dynstack.stack.length
			newLength   = oldLength * 2
			newElements = make([]T, newLength)
		)
		for i := uint32(0); i < oldLength; i++ {
			newElements[i] = dynstack.stack.elements[i]
		}
		dynstack.stack.elements = newElements
		dynstack.stack.length = newLength
	}
	return dynstack.stack.Push(item)
}

func (dynstack *DynamicStack[T]) full() bool {
	return dynstack.stack.Full()
}

func (dynstack *DynamicStack[T]) Peek() (item T, ok bool) {
	return dynstack.stack.Peek()
}

func (dynstack *DynamicStack[T]) Pop() (item T, ok bool) {
	return dynstack.stack.Pop()
}

func (dynstack *DynamicStack[T]) Empty() bool {
	return dynstack.stack.Empty()
}

func (dynstack *DynamicStack[T]) Length() int64 {
	return dynstack.stack.Length()
}
