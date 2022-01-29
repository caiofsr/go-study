package concurrency

import (
	"reflect"
	"testing"
)

func mockVerifierWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func TestVerifyWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	result := VerifyWebsites(mockVerifierWebsite, websites)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}
