package kata


func Gimme(array [3]int) int {
	minIndex, maxIndex := 0, 0
	for i, num := range array {
		if num < array[minIndex] {
			minIndex = i
		}
		if num > array[maxIndex] {
			maxIndex = i
		}
	}
	var resultIndex int
	for i := range array {
		if i != minIndex && i != maxIndex {
			resultIndex = i
			break
		}
	}
	return resultIndex
}
