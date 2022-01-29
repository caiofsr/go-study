package concurrency

import "time"

type VerifierWebsite func(string) bool
type result struct {
	string
	bool
}

func VerifyWebsites(vw VerifierWebsite, urls []string) map[string]bool {
	results := make(map[string]bool)
	channelResult := make(chan result)

	for _, url := range urls {
		go func(u string) {
			channelResult <- result{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channelResult
		results[result.string] = result.bool
	}

	time.Sleep(2 * time.Second)

	return results
}
