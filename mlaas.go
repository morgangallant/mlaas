// Package mlaas allocates 100MiB every hour, without freeing it.
package mlaas

import (
	"fmt"
	"time"
)

const allocMiB = 100

// Start begins leaking memory every hour. It never returns, and
// we deliberately don't accept a context or any sort of shutdown
// signal to make it easy to shut down or de-allocate the memory.
func Start() {
	var buffers [][]byte
	for {
		time.Sleep(time.Hour)
		buf := make([]byte, allocMiB<<20)
		for i := range buf {
			if i%2 == 0 {
				buf[i] = 0x69
			} else {
				buf[i] = 0x42
			}
		}
		buffers = append(buffers, buf)
		fmt.Printf("[mlaas] total memory allocated: %dMiB\n", len(buffers)*allocMiB)
	}
}
