//1. Иногда приходят нули. В чем проблема? Исправь ее
//2. Если функция bank_network_call выполняется 5 секунд,
//то за сколько выполнится balance()? Как исправить проблему?
//3. Представим, что bank_network_call возвращает ошибку дополнительно.
//Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.


func balance() int {
	x := make(map[int]int, 1)
	var m sync.Mutex

	// call bank
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			m.Lock()
			b := bank_network_call(i)
			x[i] = b
			m.Unlock()
		}()
	}

	// Как-то считается сумма значений в мапе и возвращается
	return sumOfMap
}


//Ответ:

func balance() (int, error) {
	x := make(map[int]int, 5)

	var (
		mu sync.Mutex
		wg sync.WaitGroup
		sOnce sync.Once
		errReturn error
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			defer wg.Done()
			b, err := bank_network_call(i)
			if err != nil {
				sOnce.Do(func() {
					errReturn = err
					cancel() // Отменяем все остальные вызовы
				})
			return
		}

			select {
				case <-ctx.Done(): // Если другая горутина уже вызвала cancel()
				return

				default:
			}

			mu.Lock()
			x[i] = b
			mu.Unlock()
		}()
	}

	wg.Wait()
	if errReturn != nil {
		return 0, errReturn
	}

// Как-то считается сумма значений в мапе и возвращается
	return sumOfMap, nil
}