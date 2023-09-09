package iter

// Accumulate accumulates the values of the iterator based on the provided function.
func Accumulate[E any](it Iter[E], f func(e1, e2 E) E, init E) Iter[E] {
	return IterFunc[E](func() (E, bool) {
		e, ok := it.Next()
		if !ok {
			return zero[E](), false
		}
		init = f(init, e)
		return init, true
	})
}

// Reduce reduces the elements of the iterator to a single value based on the provided function.
func Reduce[E any](it Iter[E], f func(e1, e2 E) E, init E) E {
	for {
		elem, ok := it.Next()
		if !ok {
			return init
		}
		init = f(init, elem)
	}
}

func next[E any](els ...E) (E, []E) {
	if len(els) == 0 {
		return zero[E](), nil
	}
	return els[0], els[1:]
}

// Chain chains multiple iterators into one.
func Chain[E any](it ...Iter[E]) Iter[E] {
	curr, it := next(it...)
	return IterFunc[E](func() (E, bool) {
		if curr == nil {
			return zero[E](), false
		}
		for {
			if curr == nil {
				return zero[E](), false
			}
			e, ok := curr.Next()
			if !ok {
				curr, it = next(it...)
				if curr == nil {
					return zero[E](), false
				}
				continue
			}
			return e, true
		}
	})
}

// DropWhile drops elements from the iterator while the provided function returns true.
func DropWhile[E any](pred func(E) bool, it Iter[E]) Iter[E] {
	var droppedAll bool
	return IterFunc[E](func() (E, bool) {
		for {
			e, ok := it.Next()
			if !ok {
				return zero[E](), false
			}
			if !droppedAll {
				if pred(e) {
					continue
				}
				droppedAll = true
			}
			return e, true
		}
	})
}

// FilterFalse filters out the elements for which the provided function returns true.
func FilterFalse[E any](pred func(E) bool, it Iter[E]) Iter[E] {
	return IterFunc[E](func() (E, bool) {
		for {
			e, ok := it.Next()
			if !ok {
				return zero[E](), false
			}
			if !pred(e) {
				return e, true
			}
		}
	})
}

// TakeWhile takes elements from the iterator while the provided function returns true.
func TakeWhile[E any](pred func(E) bool, it Iter[E]) Iter[E] {
	var taken bool
	return IterFunc[E](func() (E, bool) {
		if taken {
			return zero[E](), false
		}
		e, ok := it.Next()
		if !ok {
			return zero[E](), false
		}
		if pred(e) {
			return e, true
		}
		taken = true
		return zero[E](), false
	})
}
