// Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Modified by Andy Dalton to implement exercise solution.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Fetchall fetches URLs in parallel and reports their times and sizes.
func main() {
	start := time.Now()
	ch := make(chan string)

	for index, url := range os.Args[1:] {
		go fetch(start.UnixMilli(), index, url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(ts int64, index int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// Open for write, create if not exist, truncate if exist
	f, err := os.OpenFile(fmt.Sprintf("output_%v_%v.txt", ts, index), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer f.Close()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
