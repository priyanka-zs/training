package main

import "fmt"

func fib() {
	n := 10
	f1 := 0
	f2 := 1
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			fmt.Println(i)
		} else {
			f := f1 + f2
			fmt.Println(f)
			f1 = f2
			f2 = f
		}
	}
}
