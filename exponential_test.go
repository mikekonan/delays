package delays

import (
	"fmt"
	"testing"
	"time"
)

// TestExponential tests the Exponential function and generateDelays method.
func TestExponential(t *testing.T) {
	tests := []struct {
		name           string
		totalDuration  time.Duration
		numAttempts    int
		exponent       float64
		expectedDelays []time.Duration
	}{
		{
			name:           "Basic exponential",
			totalDuration:  30 * time.Second,
			numAttempts:    10,
			exponent:       2.0,
			expectedDelays: []time.Duration{0, 0, 0, 1 * time.Second, 2 * time.Second, 3 * time.Second, 4 * time.Second, 5 * time.Second, 7 * time.Second, 9 * time.Second},
		},
		{
			name:           "Single attempt",
			totalDuration:  30 * time.Second,
			numAttempts:    1,
			exponent:       2.0,
			expectedDelays: []time.Duration{30 * time.Second},
		},
		{
			name:           "Zero exponent",
			totalDuration:  30 * time.Second,
			numAttempts:    10,
			exponent:       0.0,
			expectedDelays: []time.Duration{3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second, 3 * time.Second},
		},
		{
			name:           "Negative exponent",
			totalDuration:  30 * time.Second,
			numAttempts:    10,
			exponent:       -2.0,
			expectedDelays: []time.Duration{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strategy := Exponential(tt.totalDuration, tt.numAttempts, tt.exponent)
			delays := strategy.Plan()

			if len(delays) != len(tt.expectedDelays) {
				t.Errorf("%s: expected %d delays, got %d", tt.name, len(tt.expectedDelays), len(delays))
			}

			for i, delay := range delays {
				if delay != tt.expectedDelays[i] {
					t.Errorf("%s: expected delay at index %d to be %v, got %v", tt.name, i, tt.expectedDelays[i], delay)
				}
			}
		})
	}
}

// TestAt tests the At method of the ExponentialStrategy struct.
func TestAt(t *testing.T) {
	strategy := Exponential(30*time.Second, 10, 2.0)

	tests := []struct {
		name    string
		attempt int
		want    time.Duration
	}{
		{"First attempt", 0, 0 * time.Second},
		{"Last attempt", 9, 9 * time.Second},
		{"Out of bounds negative", -1, 0 * time.Second},
		{"Out of bounds positive", 10, 0 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strategy.At(tt.attempt)
			if got != tt.want {
				t.Errorf("%s: At(%d) = %v; want %v", tt.name, tt.attempt, got, tt.want)
			}
		})
	}
}

// TestString tests the String method of the ExponentialStrategy struct.
func TestString(t *testing.T) {
	strategy := Exponential(30*time.Second, 10, 2.0)

	want := "Delays:\n"
	expectedDelays := strategy.Plan()
	for i, delay := range expectedDelays {
		want += fmt.Sprintf("Attempt %d: %v\n", i+1, delay)
	}

	t.Run("String output", func(t *testing.T) {
		if got := strategy.String(); got != want {
			t.Errorf("String output: expected %v, got %v", want, got)
		}
	})
}
