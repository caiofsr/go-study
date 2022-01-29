package concurrency

import (
	"testing"
	"time"
)

func slowStubVerifierWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func BenchmarkVerifyWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}

	for i := 0; i < b.N; i++ {
		VerifyWebsites(slowStubVerifierWebsite, urls)
	}
}
