package main

import (
	"fmt"
	"time"

	"github.com/jonioliveira/test-log/pkg/injection"
)

func main() {
	start := time.Now()

	startLog := time.Now()
	log := injection.New(injection.DebugLevel)
	fmt.Printf("time to initialize log %v\n", time.Since(startLog))

	startfn := time.Now()
	injection.Execute(log, "+", 2, 3)
	fmt.Printf("time to execute sum is %v\n", time.Since(startfn))

	startfn = time.Now()
	injection.Execute(log, "-", 5, 4)
	fmt.Printf("time to execute sub is %v\n", time.Since(startfn))

	startfn = time.Now()
	injection.Execute(log, "+")
	fmt.Printf("time to execute sub is %v\n", time.Since(startfn))

	startfn = time.Now()
	injection.Execute(log, "-", 4, 5)
	fmt.Printf("time to execute sub is %v\n", time.Since(startfn))

	end := time.Since(start)
	log.Debugf("total time is %v", end)
}
