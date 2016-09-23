package bootstrap

import "time"

var startTime float64

// Reset measurement of elapsed time
func Reset() {
	startTime = 0
}

// Now measures how long an operation took thats over 0.000001s
func Now() float64 {
	myTime := float64(time.Now().UnixNano()) / 1000000.0
	if startTime < 0.000001 {
		startTime = myTime
	}

	return myTime - startTime
}
