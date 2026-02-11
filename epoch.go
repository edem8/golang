// The unic epoch- the number of non leap seconds that have elapsed since
// January 1, 1970 Thursday at 00:00:00 UTC

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main(){
	// Get the current time
	now := time.Now()
	p(now)

	// get the current unix epoch
	p(now.Unix())
	p(now.UnixNano()) //get it in Nanoseconds
	p(now.UnixMilli()) // get it in milliseconds


	// converting the integer seconds or nanoseconds 
	// since the epoch back to time
	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))
}