type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
// Нужно реализовать функцию, которая выполняет
// поиск query во всех переданных SearchFunc
// Когда получаем первый успешный результат -
// отдаем его сразу. Если все SearchFunc отработали
// с ошибкой - отдаем последнюю полученную ошибку
}


//Ответ:

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {

	if len(sfs) == 0 {
		return Result{}, errors.New("no search functions provided")
	}
  
	resultChan := make(chan Result, 1)
	errChan := make(chan error, len(sfs))

	cCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, sf := range sfs {
		go func(sf SearchFunc) {
			// Если контекст уже отменён — выходим
			if cCtx.Err() != nil {
				return
			}

			res, err := sf(cCtx, query)
			if err != nil {
				select {
					case errChan <- err:
					case <-cCtx.Done():
				}
				return
			}

// Первый успешный ответ отправляем в resultChan и cancel()
			select {
				case resultChan <- res:
				cancel() // Закрываем контекст → остальные горутины прервутся
				case <-cCtx.Done():
			}
		}(sf)
	}

	// Слушаем, пока не придёт первый успешный результат или не переберём все ошибки
	var lastErr error
	for i := 0; i < len(sfs); i++ {
		select {
			case res := <-resultChan:
			return res, nil // Первый успех
			
			case err := <-errChan:
			lastErr = err

			case <-ctx.Done(): // Если глобальный контекст был отменён
			return Result{}, ctx.Err()
		}
	}

	return Result{}, lastErr
}