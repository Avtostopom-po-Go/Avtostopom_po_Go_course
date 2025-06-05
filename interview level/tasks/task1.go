//Условие задачи
//Дан массив целых чисел nums и целое число k. Нужно написать функцию,
//которая вынимает из массива nums k наиболее часто встречающихся элементов.

//Пример
//# ввод
//nums = [1,1,1,2,2,3]
//k = 2
//# вывод (в любом порядке)
//[1, 2]


func topKFrequentElements(nums []int, k int) []int {
    count := make(map[int]int)
    freq := make([][]int, len(nums) + 1)

    for _, n := range nums{
        count[n]++
    }

    for n, c := range count {
        freq[c] = append(freq[c], n)
    }
    res := make([]int, 0, k)

    for i:= len(freq); i>=0; i--{
        for _, v := range freq[i]{
            res = append(res, v)
            if len(res) == k{
                return res
            }
        }
    }
    return res
}