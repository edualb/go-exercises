// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var path = "./fetchall.txt"

func main() {
	start := time.Now()
	ch := make(chan string)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for range os.Args[1:] {
		_, err = file.WriteString(<-ch) // receive from channel ch to fetchall.txt file
		if err != nil {
			fmt.Println(err)
		}
	}
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}

//!-
