package iter

// Count returns an infinite iterator starting from the
// given value and incrementing by the specified step.
func Count(start, step int) Iter[int] {
	return IterFunc[int](func() (int, bool) {
		res := start
		start += step
		return res, true
	})
}

// Cycle cycles through the elements of the iterator indefinitely.
func Cycle[E any](it Iter[E]) Iter[E] {
	var exausted bool
	var el []E
	var i int
	return IterFunc[E](func() (E, bool) {
	exausted:
		if exausted {
			c := i
			i = (i + 1) % len(el)
			return el[c], true
		}
		e, ok := it.Next()
		if ok {
			el = append(el, e)
			return e, true
		}
		exausted = true
		if len(el) == 0 {
			panic("iterable was empty")
		}
		goto exausted
	})
}

// Repeat repeats the given element for the specified number of times.
// If <times> is a negative number, the element is repeated indefinitely.
func Repeat[E any](e E, times int) Iter[E] {
	var occ int
	return IterFunc[E](func() (E, bool) {
		if times < 0 {
			return e, true
		}
		if occ < times {
			occ++
			return e, true
		}
		return zero[E](), false
	})
}
