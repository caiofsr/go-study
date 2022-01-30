package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsLimit = 10 * time.Second

func configurableRunner(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time limit exceeded for %s and %s", urlA, urlB)
	}
}

func Runner(urlA, urlB string) (winner string, err error) {
	return configurableRunner(urlA, urlB, tenSecondsLimit)
}

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
