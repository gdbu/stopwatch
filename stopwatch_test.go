package stopwatch

import (
	"fmt"
	"testing"
	"time"
)

func TestStopStart(t *testing.T) {
	var s Stopwatch
	s.Start()
	time.Sleep(time.Second * 3)
	fmt.Printf("Slept for %v\n", s.Stop())
}
