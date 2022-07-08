package creek

import "time"

// The Wait function pauses the current stream for the duration passed.
// The first and only parameter expects a value from the built-in time.Duration package.
func (s Stream[T]) Wait(duration time.Duration) Stream[T] {
	time.Sleep(duration)
	return s
}
