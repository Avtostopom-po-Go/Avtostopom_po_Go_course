//1. Merge n channels
//2. Если один из входных каналов закрывается,
//то нужно закрыть все остальные каналы

func case3(channels ...chan int) chan int {


//Ответ:

func case3(channels ...chan int) chan int {
	out := make(chan int) // Выходной канал

	ctx, cancel := context.WithCancel(context.Background()) // Контекст для отмены

	var wg sync.WaitGroup

// Запускаем горутину для каждого входного канала
	for _, ch := range channels {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for {
				// Если уже отменено, сразу выходим
				select {
					case <-ctx.Done():
					return
				
					default:
				}

			// Читаем данные из канала
				val, ok := <-c
				if !ok { // Если канал закрыт, останавливаем всё
					cancel()
					return
				}
				out <- val
			}
		}(ch)
	}

// Горутина для закрытия out после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}