// this provide a simple overview about CPU usage

package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func cpu(start time.Time, done chan struct{}) {
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10000; i++ {
		time.Sleep(time.Millisecond)
	}

	// this part informs the "done" channel that CPU profiling is completed 
	done <- struct{}{}
}

func main() {
	start := time.Now()

	done := make(chan struct{})

	go cpu(start, done)

	<-done

	elapsed := time.Since(start)
	fmt.Printf("page took %s\n", elapsed)
}
// to finish this script
