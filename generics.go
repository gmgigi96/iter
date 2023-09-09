package iter

// IterFunc is a type that represents an iterator function.
type IterFunc[E any] func() (E, bool)

func (f IterFunc[E]) Next() (E, bool) {
	return f()
}

// Filter filters the elements of the iterator based on the provided function.
func Filter[E any](f func(E) bool, it Iter[E]) Iter[E] {
	return IterFunc[E](func() (E, bool) {
		for {
			e, ok := it.Next()
			if !ok {
				return zero[E](), false
			}
			if f(e) {
				return e, true
			}
		}
	})
}

// Map maps the elements of the iterator to another type based on the provided function.
func Map[E, T any](f func(E) T, it Iter[E]) Iter[T] {
	return IterFunc[T](func() (T, bool) {
		e, ok := it.Next()
		if !ok {
			return zero[T](), false
		}
		return f(e), true
	})
}

// Range returns an iterator for a range of integers up to the given stop.
func Range(stop int) Iter[int] {
	return Range2(0, stop)
}

// Range2 returns an iterator for a range of integers between start and stop.
func Range2(start, stop int) Iter[int] {
	return Range3(start, stop, 1)
}

// Range3 returns an iterator for a range of integers
// between start and stop with a specified step.
func Range3(start, stop, step int) Iter[int] {
	if step == 0 {
		panic("step cannot be zero")
	}
	return IterFunc[int](func() (int, bool) {
		if step > 0 && start >= stop || step < 0 && start <= stop {
			return 0, false
		}
		curr := start
		start += step
		return curr, true
	})
}

// Enum represents a value with its index.
type Enum[E any] struct {
	Index int
	Value E
}

// Enumerate enumerates the elements of the iterator.
func Enumerate[E any](it Iter[E]) Iter[Enum[E]] {
	var i int
	return IterFunc[Enum[E]](func() (Enum[E], bool) {
		e, ok := it.Next()
		if !ok {
			return zero[Enum[E]](), false
		}
		n := i
		i++
		return Enum[E]{Value: e, Index: n}, true
	})
}

// Zip zips two iterators into one.
func Zip[E, T any](it1 Iter[E], it2 Iter[T]) Iter[Pair[E, T]] {
	return IterFunc[Pair[E, T]](func() (Pair[E, T], bool) {
		first, ok := it1.Next()
		if !ok {
			return zero[Pair[E, T]](), false
		}
		second, ok := it2.Next()
		if !ok {
			return zero[Pair[E, T]](), false
		}
		return Pair[E, T]{First: first, Second: second}, true
	})
}
