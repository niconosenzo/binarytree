package binary

//Search will receive a slice of integers along with a number to search within the slice
func Search(guess int, list []int) bool {

	low := 0
	high := len(list) - 1

	for low <= high {
		median := (low + high) / 2

		if list[median] < guess {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(list) || list[low] != guess {
		return false
	}

	return true
}

// func main() {
// 	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
// 	fmt.Println(binarySearch(63, items))
// }
