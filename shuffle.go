package creek

import "math/rand"

// The Shuffle function shuffles the stream.
func (s Stream[T]) Shuffle() Stream[T] {
	rand.Shuffle(len(s.Array), func(i, j int) {
		s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
	})

	return s
}

// The Shuffle function shuffles the stream.
func (s StructStream[T]) Shuffle() StructStream[T] {
	rand.Shuffle(len(s.Array), func(i, j int) {
		s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
	})

	return s
}

// The Shuffle function shuffles the stream.
func (s MapStream[T, V]) Shuffle() MapStream[T, V] {
	rand.Shuffle(len(s.Array), func(i, j int) {
		s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
	})

	return s
}
