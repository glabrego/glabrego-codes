package main

import "fmt"
import "time"

func main() {
	puts := fmt.Println

	now := time.Now()
	puts(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	puts(then)

	puts(then.Year())
	puts(then.Month())
	puts(then.Day())
	puts(then.Hour())
	puts(then.Minute())
	puts(then.Second())
	puts(then.Nanosecond())
	puts(then.Location())

	puts(then.Weekday())

	puts(then.Before(now))
	puts(then.After(now))
	puts(then.Equal(now))

	diff := now.Sub(then)
	puts(diff)

	puts(diff.Hours())
	puts(diff.Minutes())
	puts(diff.Seconds())
	puts(diff.Nanoseconds())

	puts(then.Add(diff))
	puts(then.Add(-diff))
}
