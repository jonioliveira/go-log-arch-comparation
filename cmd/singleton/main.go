package main

import (
	"fmt"
	"time"

	"github.com/jonioliveira/test-log/pkg/singleton"
)

func main() {
	start := time.Now()

	startfn := time.Now()
	singleton.Execute("+", 2, 3)
	fmt.Printf("time to execute sum is %v\n", time.Since(startfn))

	startfn = time.Now()
	singleton.Execute("-", 5, 4)
	fmt.Printf("time to execute sub is %v\n", time.Since(startfn))

	startfn = time.Now()
	singleton.Execute("+")
	fmt.Printf("time to execute sub is %v\n", time.Since(startfn))

	startfn = time.Now()
	singleton.Execute("-", 4, 5)
	fmt.Printf("time to execute sub is %v\n", time.Since(startfn))

	end := time.Since(start)
	singleton.GetInstance().Debugf("total time is %v", end)
}
