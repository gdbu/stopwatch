package stopwatch

import "time"

// Stopwatch represents a stopwatch
type Stopwatch struct {
	start time.Time
}

// Start will begin the current stopwatch
func (s *Stopwatch) Start() {
	s.start = time.Now()
}

// Stop will stop the current stopwatch
func (s *Stopwatch) Stop() (duration time.Duration) {
	end := time.Now()
	return time.Duration(end.UnixNano() - s.start.UnixNano())
}
