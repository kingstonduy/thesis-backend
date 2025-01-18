package math_utils

import (
	"time"
)

var fibonacciNumbers = []int{
	1,       // F(0)
	1,       // F(1)
	2,       // F(2)
	3,       // F(3)
	5,       // F(4)
	8,       // F(5)
	13,      // F(6)
	21,      // F(7)
	34,      // F(8)
	55,      // F(9)
	89,      // F(10)
	144,     // F(11)
	233,     // F(12)
	377,     // F(13)
	610,     // F(14)
	987,     // F(15)
	1597,    // F(16)
	2584,    // F(17)
	4181,    // F(18)
	6765,    // F(19)
	10946,   // F(20)
	17711,   // F(21)
	28657,   // F(22)
	46368,   // F(23)
	75025,   // F(24)
	121393,  // F(25)
	196418,  // F(26)
	317811,  // F(27)
	514229,  // F(28)
	832040,  // F(29)
	1346269, // F(30)
	2178309, // F(31)
}

// GetNextTimeExecution returns the time after adding `n` seconds to `now`,
// where `n` is a power of 2.
func GetNextTimeExecution(n int, now time.Time) time.Time {
	return time.Now()

	// n = int(math.Max(float64(n), 0))
	// n = int(math.Min(float64(n), float64(len(fibonacciNumbers)-1)))

	// duration := time.Duration(fibonacciNumbers[n]) * time.Minute
	// return now.Add(duration)
}
