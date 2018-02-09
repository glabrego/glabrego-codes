package main

import "fmt"
import "time"

func main() {
	puts := fmt.Println

	now := time.Now()
	puts(now.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	puts(t1)

	puts(now.Format("3:04PM"))
	puts(now.Format("Mon Jan _2 15:04:05 2006"))
	puts(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	puts(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	puts(e)
}
