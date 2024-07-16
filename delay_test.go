package delays

import (
	"testing"
	"time"
)

// TestNew tests the New function.
func TestNew(t *testing.T) {
	strategy := Exponential(30*time.Second, 10, 2.0)
	delay := New(strategy)

	t.Run("New Delay Initialization", func(t *testing.T) {
		if delay.strategy != strategy {
			t.Errorf("New Delay Initialization: expected strategy %v, got %v", strategy, delay.strategy)
		}
	})
}

// TestDelayAt tests the At method of the Delay struct.
func TestDelayAt(t *testing.T) {
	strategy := Exponential(30*time.Second, 10, 2.0)
	delay := New(strategy)

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
			got := delay.At(tt.attempt)
			if got != tt.want {
				t.Errorf("%s: At(%d) = %v; want %v", tt.name, tt.attempt, got, tt.want)
			}
		})
	}
}

// TestDelayPlan tests the Plan method of the Delay struct.
func TestDelayPlan(t *testing.T) {
	strategy := Exponential(30*time.Second, 10, 2.0)
	delay := New(strategy)

	t.Run("Delay Plan", func(t *testing.T) {
		want := strategy.Plan()
		got := delay.Plan()

		if len(got) != len(want) {
			t.Errorf("Delay Plan: expected %d delays, got %d", len(want), len(got))
		}

		for i, delay := range got {
			if delay != want[i] {
				t.Errorf("Delay Plan: plan delay at index %d = %v; want %v", i, delay, want[i])
			}
		}
	})
}

// TestDelayString tests the String method of the Delay struct.
func TestDelayString(t *testing.T) {
	strategy := Exponential(30*time.Second, 10, 2.0)
	delay := New(strategy)

	t.Run("Delay String", func(t *testing.T) {
		want := strategy.String()
		if got := delay.String(); got != want {
			t.Errorf("Delay String: expected %v, got %v", want, got)
		}
	})
}
