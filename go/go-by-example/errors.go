package main

import (
	"errors"
	"fmt"
)

func function1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func function2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := function1(i); e != nil {
			fmt.Println("function1 failed:", e)
		} else {
			fmt.Println("function1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := function2(i); e != nil {
			fmt.Println("function2 failed:", e)
		} else {
			fmt.Println("function2 worked:", r)
		}
	}

	_, e := function2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
