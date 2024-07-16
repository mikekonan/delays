package delays

import (
	"fmt"
	"math"
	"time"
)

// ExponentialStrategy implements the DelayStrategy interface.
type ExponentialStrategy struct {
	totalDuration time.Duration
	numAttempts   int
	exponent      float64
	delays        []time.Duration
}

// Exponential is a public constructor for ExponentialStrategy.
func Exponential(totalDuration time.Duration, numAttempts int, exponent float64) *ExponentialStrategy {
	expDelay := &ExponentialStrategy{
		totalDuration: totalDuration,
		numAttempts:   numAttempts,
		exponent:      exponent,
	}

	expDelay.generateDelays()

	return expDelay
}

func (e *ExponentialStrategy) generateDelays() {
	if e.exponent < 0 || e.totalDuration <= 0 || e.numAttempts <= 0 {
		return
	}

	e.delays = make([]time.Duration, e.numAttempts)
	if e.numAttempts == 1 {
		e.delays[0] = e.totalDuration
		return
	}

	// Generate exponentially distributed values between 0 and 1
	expValues := make([]float64, e.numAttempts)
	for i := 0; i < e.numAttempts; i++ {
		expValues[i] = math.Pow(float64(i)/float64(e.numAttempts-1), e.exponent)
	}

	// Normalize the exponential values to sum to 1
	sumExpValues := 0.0
	for _, val := range expValues {
		sumExpValues += val
	}

	normalizedExpValues := make([]float64, e.numAttempts)
	for i, val := range expValues {
		normalizedExpValues[i] = val / sumExpValues
	}

	// Scale the normalized values to fit within the total_duration
	totalDurationSeconds := e.totalDuration.Seconds()
	scaledDelays := make([]float64, e.numAttempts)
	for i, val := range normalizedExpValues {
		scaledDelays[i] = val * totalDurationSeconds
	}

	// Calculate the individual delays
	for i, delay := range scaledDelays {
		e.delays[i] = time.Duration(math.Round(delay)) * time.Second
	}
}

// At returns the delay for a specific attempt.
func (e *ExponentialStrategy) At(attempt int) time.Duration {
	if attempt < 0 || attempt >= e.numAttempts {
		return 0
	}

	return e.delays[attempt]
}

// Plan returns the full plan of delays.
func (e *ExponentialStrategy) Plan() []time.Duration {
	return e.delays
}

// String returns the plan of delays as a formatted string.
func (e *ExponentialStrategy) String() string {
	result := "Delays:\n"
	for i, delay := range e.delays {
		result += fmt.Sprintf("Attempt %d: %v\n", i+1, delay)
	}
	return result
}
