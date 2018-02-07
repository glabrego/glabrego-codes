package main

import s "strings"
import "fmt"

var puts = fmt.Println

func main() {
	puts("Contains:  ", s.Contains("test", "es"))
	puts("Count:     ", s.Count("test", "t"))
	puts("HasPrefix: ", s.HasPrefix("test", "te"))
	puts("HasSuffix: ", s.HasSuffix("test", "st"))
	puts("Index:     ", s.Index("test", "e"))
	puts("Join:      ", s.Join([]string{"a", "b"}, "-"))
	puts("Repeat:    ", s.Repeat("a", 5))
	puts("Replace:   ", s.Replace("foo", "o", "0", -1))
	puts("Replace:   ", s.Replace("foo", "o", "0", 1))
	puts("Split:     ", s.Split("a-b-c-d-e", "-"))
	puts("ToLower:   ", s.ToLower("TEST"))
	puts("ToUpper:   ", s.ToUpper("test"))
	puts()

	puts("Len: ", len("hello"))
	puts("Char:", "hello"[1])
}
