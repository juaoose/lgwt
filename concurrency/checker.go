package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// Go has a built in race detector
// go test -race
// https://go.dev/blog/race-detector

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		/* Anonymous functions can be executed at teh same
		time as they are declared and they maintain access to
		the lexical scope they are defined in
		*/

		go func(u string) {
			// Send statement (channel <- value)
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive statement (value <- channel)
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
