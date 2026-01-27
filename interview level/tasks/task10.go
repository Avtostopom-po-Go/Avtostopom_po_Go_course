//1. Merge n channels
//2. Если один из входных каналов закрывается,
//то нужно закрыть все остальные каналы

func case3(channels ...chan int) chan int {
