package main

import (
	"flag"
	"fmt"
	"github.com/mikekonan/delays"
	"math"
	"time"
)

func main() {
	strategy := flag.String("strategy", "exponential", "Delay strategy (exponential, linear)")
	totalDuration := flag.Duration("totalDuration", 2*time.Minute, "Total duration")
	numAttempts := flag.Int("numAttempts", 10, "Number of attempts")
	exponent := flag.Float64("exponent", math.SqrtPhi, "Exponent for the exponential strategy (only used in exponential strategy)")

	flag.Parse()

	var delayStrategy delays.DelayStrategy
	switch *strategy {
	case "exponential":
		delayStrategy = delays.Exponential(*totalDuration, *numAttempts, *exponent)
	default:
		fmt.Println("Unknown strategy:", *strategy)
		return
	}

	delay := delays.New(delayStrategy)

	printPlanWithElapsed(delay.Plan())
}

func printPlanWithElapsed(plan []time.Duration) {
	fmt.Println("Plan of Delays:")
	fmt.Println("================")
	elapsed := time.Duration(0)
	for i, delay := range plan {
		elapsed += delay
		fmt.Printf("Attempt %2d: wait %v. Elapsed %v.\n", i+1, delay, elapsed)
	}

	fmt.Println("================")
}
