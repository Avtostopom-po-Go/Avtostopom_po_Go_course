//Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
//Её тело нельзя изменять (представим, что внутри сетевой запрос).
//Нужно написать обертку predictableFunc, которая будет работать с заданным фиксированным таймаутом (например, 1 секунду).
//Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
//Если "длинная" функция отработала за это время - отлично, возвращаем результат.
//Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
//Сигнатуру функцию обёртки менять можно.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

func predictableFunc() int64 {

}  

func main() {
	fmt.Println("started")
	fmt.Println(predictableFunc())
}


//Ответ:

package main

import (
	"fmt"
	"math/rand"
	"time"
	"context"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

func predictableFunc() (int64, error) {
	start := time.Now() // Засекаем время выполнения
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resultCh := make(chan int64, 1)

	// Запускаем unpredictableFunc в горутине

	go func() {
		resultCh <- unpredictableFunc()
	}()
  
	select {
		case res := <-resultCh:
		fmt.Println("Execution time:", time.Since(start))
		return res, nil

		case <-ctx.Done():
		fmt.Println("Execution time (timeout):", time.Since(start))

		return 0, fmt.Errorf("timeout: function took too long")
	}
}

func main() {
	fmt.Println("started")
	fmt.Println(predictableFunc())
}