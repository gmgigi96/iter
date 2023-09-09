package iter

// Iter is an interface representing an iterator.
type Iter[E any] interface {
	// Next returns the next element in the iterator.
	Next() (elem E, ok bool)
}

// Pair represents a pair of values.
type Pair[E, T any] struct {
	First  E
	Second T
}

// Slice converts the iterator to a slice.
func Slice[E any](it Iter[E]) []E {
	s := make([]E, 0)
	for e, ok := it.Next(); ok; e, ok = it.Next() {
		s = append(s, e)
	}
	return s
}

// Chan converts the iterator to an unbuffered channel.
func Chan[E any](it Iter[E]) <-chan E {
	c := make(chan E)
	go func() {
		for e, ok := it.Next(); ok; e, ok = it.Next() {
			c <- e
		}
		close(c)
	}()
	return c
}

// ChanBuff converts the iterator to a buffered channel.
func ChanBuff[E any](it Iter[E], l int) <-chan E {
	c := make(chan E, l)
	go func() {
		for e, ok := it.Next(); ok; e, ok = it.Next() {
			c <- e
		}
		close(c)
	}()
	return c
}

// FromSlice creates an iterator from a slice.
func FromSlice[E any](s []E) Iter[E] {
	var i int
	return IterFunc[E](func() (E, bool) {
		if i < len(s) {
			ix := i
			i++
			return s[ix], true
		}
		return zero[E](), false
	})
}
