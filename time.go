package main
import (
	"fmt"
	"time"
)

// aliasing the print
var p = fmt.Println


func main(){
	// Getting the current time
	now := time.Now()
	p(now)

	// We can build a time struct by providing the month, year, day, etc
	// Times are always associated with a time Location. i.e time zone
	then := time.Date(2025, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Extracting the various components of time
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Weekday extraction
	p(then.Weekday())

	// These methods compare two times. Whether they occur at the same time
	// whether one time is before or after another time
	p(then.Equal(now))
	p(then.Before(now))
	p(then.After(now))

	// sub methods return a duration representing the intervals between two times
	diff := now.Sub(then)
	p(diff)

	// computing the duration in different units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// We can use Add() to advance time by a duration
	// or use Add() with "-" to move time backward

	p(then.Add(diff))
	p(now.Add(-diff))

}