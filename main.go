// SPDX-License-Identifier: BSD-3-Clause
// random creates a random number

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

const MaxInt = math.MaxInt64

var (
	upper int
	lower int
)

func main() {
	flag.IntVar(&upper, "upper", MaxInt, "Upper limit for random number")
	flag.IntVar(&upper, "u", MaxInt, "Upper limit for random number (shorthand)")
	flag.IntVar(&lower, "lower", 0, "Lower limit for random number")
	flag.IntVar(&lower, "l", 0, "Lower limit for random number (shorthand)")
	flag.Parse()

	if upper < lower {
		log.Fatal("Upper limit must be greater than or equal to lower limit")
	}

	if upper == lower {
		log.Fatal("Upper limit must be greater than lower limit")
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(upper-lower) + lower
	fmt.Println(randomNumber)
}
