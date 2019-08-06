package diagonalsum

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	numRows := len(arr)

	startRow := 0
	leftIndex := -1
	rightIndex := -1
	var leftSum int32
	var rightSum int32
	for i := 0; i < numRows; i++ {
		if len(arr[i]) == (numRows-i) && leftIndex == -1 {
			startRow = i
			leftIndex = 0
			rightIndex = leftIndex + len(arr[i]) - 1
		}
	}

	for i := startRow; i < numRows; i++ {
		leftSum += arr[i][leftIndex]
		rightSum += arr[i][rightIndex]
		leftIndex += 1
		rightIndex -= 1
	}

	diff := leftSum - rightSum
	if diff < 0 {
		return -diff
	}
	return diff

}

// func main() {
// 	arr := [][]int32{
// 		[]int32{3},
// 		[]int32{11, 2, 4},
// 		[]int32{4, 5, 6},
// 		[]int32{10, 8, -12},
// 	}
//
// 	result := diagonalDifference(arr)
// 	fmt.Println(result)
// }
