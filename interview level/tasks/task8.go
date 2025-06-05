//Написать код, который будет выводить
//коды ответов на HTTP-запросы по двум URL
//главная страница Google и главная страница WB.
//Запросы должны осуществляться параллельно.

//Ответ:

package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.wildberries.ru",
	}

	// Создаем основной контекст с таймаутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			checkURL(url, ctx)
		}(url)
	}
	wg.Wait()
}

func checkURL(url string, ctx context.Context) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		fmt.Printf("Ошибка создания запроса для %s: %v\n", url, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Ошибка запроса для %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("URL: %s, Код ответа: %d\n", url, resp.StatusCode)
}

//если много запросов, то как внедрить воркер пул?
