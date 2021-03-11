package src


func GetFibonacci(start,end int) ( map[int32]int32 ) {


	f := make([]int32, end + 1)
	f[0] = 0
	f[1] = 1

	// формируем фибоначчи
	for i := 2; i <= end; i += 1 {
		f[i] = f[i-1] + f[i-2]
	}

	// обрезаем до нужной длинны и упаковываем
	m := make(map[int32]int32, end - start)
	for i, u := range f[start:end] {
		m[int32(i+start)] = u
	}

	return m
}