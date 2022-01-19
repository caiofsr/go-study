package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	t.Run("print until 1 then Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SpySleeper{}

		Count(buffer, sleeperSpy)

		resultado := buffer.String()
		esperado := "3\n2\n1\nGo!"

		if resultado != esperado {
			t.Errorf("got '%s', expected '%s'", resultado, esperado)
		}

		if sleeperSpy.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4, got %d", sleeperSpy.Calls)
		}
	})

	t.Run("pause before each print", func(t *testing.T) {
		spyPrintSleep := &SpyCountOperations{}

		Count(spyPrintSleep, spyPrintSleep)

		expected := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}

		if !reflect.DeepEqual(expected, spyPrintSleep.Calls) {
			t.Errorf("wanted calls %v got %v", expected, spyPrintSleep.Calls)
		}
	})
}

func TestSleeperConfigurable(t *testing.T) {
	timeSleep := 5 * time.Second

	timeSpy := &TimeSpy{}
	sleeper := SleeperConfigurator{timeSleep, timeSpy.Sleep}
	sleeper.Sleep()

	if timeSpy.durationSleep != timeSleep {
		t.Errorf("should have slept for %v but slept for %v", timeSleep, timeSpy.durationSleep)
	}
}
