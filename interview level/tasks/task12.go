type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
// Нужно реализовать функцию, которая выполняет
// поиск query во всех переданных SearchFunc
// Когда получаем первый успешный результат -
// отдаем его сразу. Если все SearchFunc отработали
// с ошибкой - отдаем последнюю полученную ошибку
}
