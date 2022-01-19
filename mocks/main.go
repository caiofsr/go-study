package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type SleeperConfigurator struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type TimeSpy struct {
	durationSleep time.Duration
}

type SpyCountOperations struct {
	Calls []string
}

func (t *TimeSpy) Sleep(duration time.Duration) {
	t.durationSleep = duration
}

func (s *SleeperConfigurator) Sleep() {
	s.sleep(s.duration)
}

func (s *SpyCountOperations) Sleep() {
	s.Calls = append(s.Calls, pause)
}

func (s *SpyCountOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const lastWord = "Go!"
const startCount = 3
const write = "write"
const pause = "pause"

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Count(out io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()

		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()

	fmt.Fprint(out, lastWord)
}

func main() {
	sleeper := &SleeperConfigurator{1 * time.Second, time.Sleep}

	Count(os.Stdout, sleeper)
}
