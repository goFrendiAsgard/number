package arithmetic

func Add(num ...int) int {
	result := 0
	for index := range num {
		result += num[index]
	}
	return result
}
