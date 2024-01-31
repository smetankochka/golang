package main

func SliceCopy(nums []int) []int {
	ans := make([]int, len(nums)) // Создаем новый слайс нужной длины и емкости
	copy(ans, nums)               // Копируем значения из исходного слайса в новый слайс
	return ans                    // Возвращаем новый слайс
}
