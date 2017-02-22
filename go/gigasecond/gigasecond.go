// Package clause.
package gigasecond

import "time"

// Version of test harness to use
const testVersion = 4

// AddGigasecond will add 1 billion seconds to the time provided.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000 * 1000 * 1000 * time.Second)
}
