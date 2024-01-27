package main

func strStr(haystack string, needle string) int {
	// Если игла (needle) пуста, вернуть 0
	if needle == "" {
		return 0
	}

	// Проход по строке (хейстеку) с использованием окна размером len(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// Проверка совпадения подстроки (иглы) с текущим окном
		if haystack[i:i+len(needle)] == needle {
			return i // Возвращаем индекс первого вхождения
		}
	}

	return -1 // Игла не найдена в хейстеке
}

func main() {
	println(strStr("sadbutsad", "sad"))
	println(strStr("leetcode", "leeto"))
}
