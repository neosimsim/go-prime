// Copyright © 2018, Alexander Ben Nasrallah <me@abn.sh>
// Use of this source code is governed by a BSD 3-clause
// style license that can be found in the LICENSE file.

package main

import "fmt"

// Returns a channel sending integers from start to infinity, or exit which is
// more likely to be the first.
func naturals(start int) <-chan int {
	naturals := make(chan int)
	go func() {
		n := start
		for true {
			naturals <- n
			n++
		}
	}()
	return naturals
}

// Takes each int from in and writes it to the returned channel, if not dividable
// by n.
func filter(n int, in <-chan int) <-chan int {
	filtered := make(chan int)
	go func() {
		for true {
			if m := <-in; m%n != 0 {
				filtered <- m
			}
		}
	}()
	return filtered
}

func main() {
	primeChan := naturals(2)

	for true {
		nextPrime := <-primeChan
		fmt.Println(nextPrime)
		primeChan = filter(nextPrime, primeChan)
	}
}
