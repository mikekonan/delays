# Delays

Delays is a Go library for generating and printing plans of delays based on different strategies. It currently supports
exponential and linear delay strategies. Additionally, there is a command-line interface (CLI) tool for testing the
library.

## Installation

To install the library, run:

```sh
go get github.com/mikekonan/delays@latest
```

To install the CLI tool, run:

```sh
go install github.com/mikekonan/delays/cmd@latest
```

### Usage

As a Library
You can use the Delays library in your Go programs to generate plans of delays. Here are examples of how to use the
exponential and linear delay strategies:

Exponential Strategy

```go 
package main

import (
	"fmt"
	"github.com/mikekonan/delays"
	"time"
)

func main() {
	totalDuration := 20 * time.Minute
	numAttempts := 5
	exponent := 1.5

	strategy := delays.Exponential(totalDuration, numAttempts, exponent)
	delay := delays.New(strategy)

	fmt.Println(delay.String())

	// Iterate over attempts and use At()
	fmt.Println("Delays using At():")
	for i := 0; i < numAttempts; i++ {
		fmt.Printf("Attempt %d: %v\n", i+1, delay.At(i))
		time.Sleep(delay.At(i))
	}
}

```

### Using the CLI Tool

The CLI tool can be used to generate delay plans for testing purposes. Here’s how you can use it:

```sh
delaycli -strategy exponential -totalDuration 20m -numAttempts 5 -exponent 1.5
```

#### Output

```sh
Plan of Delays:
================
Attempt  1: wait 2m0s. Elapsed 0s.
Attempt  2: wait 4m0s. Elapsed 2m0s.
Attempt  3: wait 6m0s. Elapsed 6m0s.
Attempt  4: wait 8m0s. Elapsed 12m0s.
Attempt  5: wait 10m0s. Elapsed 20m0s.
================
```

### Contributing

Feel free to submit issues, fork the repository, and send pull requests. For major changes, please open an issue first
to discuss what you would like to change.

### License

This project is licensed under the MIT License.