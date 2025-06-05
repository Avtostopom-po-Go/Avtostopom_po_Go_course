func main() {
	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

	fmt.Println(callRequestsForURLs(urls, 3))
}

// дернуть N урлов с лимитом K (то есть не больше K активных запросов одновременных), сигнатура:
func callRequestsForURLs(urls []string, K int) []*http.Response {}


//Ответ:

func main() {
	urls := []string{"https://example.com", "https://example.org", "https://example.net"}
	fmt.Println(callRequestsForURLs(urls, 3))
}

func callRequestsForURLs(urls []string, K int) []*http.Response {

	response := make([]*http.Response, 0, len(urls))
	maxcall := make(chan struct{}, K)
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, val := range urls {
		val := val
		maxcall <- struct{}{}

		go func() {
			defer func() {
				<-maxcall
				wg.Done()
			}()

			client := http.Client{
				Timeout: 2 * time.Second,
			}

			req, err := client.Get(val)
			if err != nil {
				return
			}

			mu.Lock()
			response = append(response, req)
			mu.Unlock()
		}()
	}

	wg.Wait()
	return response
}