// Package gigasecond implements utility for manipuating Time
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns a given date added 1 GigaSecond (1000000000 seconds)
func AddGigasecond(t time.Time) time.Time {
  return t.Add(time.Second * 1e9)
}
