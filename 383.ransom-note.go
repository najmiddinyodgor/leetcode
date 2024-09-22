package main

func CanConstruct(ransomNote string, magazine string) bool {
	arr := make(map[rune]int)

	for _, chr := range magazine {
		arr[chr]++
	}

	for _, chr := range ransomNote {
		c, ok := arr[chr]

		if ok != true || c == 0 {
			return false
		} else {
			arr[chr]--
		}
	}

	return true
}
