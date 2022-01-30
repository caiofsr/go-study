package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("return a error if server does not respond in 10 sec", func(t *testing.T) {
		serverA := createDelayedServer(11 * time.Second)
		serverB := createDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Runner(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})

	t.Run("compare speed of server returning the fast one", func(t *testing.T) {
		slowServer := createDelayedServer(20 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		result, err := Runner(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Wanted %s but got an error %v", expected, err)
		}

		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
