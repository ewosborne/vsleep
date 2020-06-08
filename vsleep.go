package main

/*
take arg for number of seconds, either int or float of up to 1 digit precision
echo stdin to stdout after sleep
*/

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/sbwhitecap/tqdm"
	"github.com/sbwhitecap/tqdm/iterators"
)

func main() {

	// take vsleep as a float, x.y, 1-digit
	// multiply by 10 to get xy
	// xy is the integer number of 100ms sleep intervals

	if len(os.Args) == 1 {
		fmt.Println("usage: vsleep <units>, where <units> are seconds. Accurate to one decimal place.")
	}

	if len(os.Args) != 2 {
		return
	}

	sleepTime, err := strconv.ParseFloat(os.Args[1], 32) // x.y seconds
	if err != nil {
		return
	}

	sleepCount := math.Round(sleepTime * 10)
	//fmt.Println("sleepTime", sleepTime, "sleepCount", sleepCount)

	if err == nil {
		tqdm.With(iterators.Interval(0, int(sleepCount)), "", func(v interface{}) (brk bool) {
			time.Sleep(100 * time.Millisecond)
			return
		})

	} else {
		fmt.Println("ERROR", err)
	}
	return
}
