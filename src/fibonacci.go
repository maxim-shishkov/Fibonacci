package src

import (
	"errors"
	"strconv"
)

func GetFibonacci(x,y string) (map[int32]int32, error) {

	start, err := strconv.Atoi(x)
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(y)
	if err != nil {
		return nil, err
	}

	if start >= end {
		return nil, errors.New("Invalid input  x >= y")
	}

	if end < 0 {
		return nil, errors.New("Invalid input  y < 0")
	}

	if start < 0 {
		return nil, errors.New("Invalid input  x < 0")
	}



	f := make([]int32, end + 1)
	f[0] = 0
	f[1] = 1

	// формируем фибоначчи
	for i := 2; i <= end; i += 1 {
		f[i] = f[i-1] + f[i-2]
	}

	// обрезаем до нужной длинны и упаковываем
	result := make(map[int32]int32, end - start)
	for i, u := range f[start:end] {
		result[int32(i+start)] = u
	}

	return result, nil
}