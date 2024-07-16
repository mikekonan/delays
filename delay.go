package delays

import "time"

// Delay holds the common configuration for generating delays.
type Delay struct {
	strategy DelayStrategy
}

// New is a public constructor for Delay.
func New(strategy DelayStrategy) *Delay {
	return &Delay{
		strategy: strategy,
	}
}

// At returns the delay for a specific attempt using the provided strategy.
func (d *Delay) At(attempt int) time.Duration {
	return d.strategy.At(attempt)
}

// Plan returns the full plan of delays using the provided strategy.
func (d *Delay) Plan() []time.Duration {
	return d.strategy.Plan()
}

// String returns the plan of delays as a formatted string using the provided strategy.
func (d *Delay) String() string {
	return d.strategy.String()
}

// DelayStrategy defines the interface for delay strategies.
type DelayStrategy interface {
	At(attempt int) time.Duration
	Plan() []time.Duration
	String() string
}
