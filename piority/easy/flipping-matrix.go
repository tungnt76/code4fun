package easy

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func FlippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	n := len(matrix[0]) / 2
	sumUpperLeftQuadrant := int32(0)
	for i := 0; i < n; i++ {
		row := matrix[i]
		if sum(row[:n]) > sum(row[len(matrix[i])-n:]) {
			sumUpperLeftQuadrant += sum(row[:n])
		} else {
			sumUpperLeftQuadrant += sum(row[len(matrix[i])-n:])
		}
	}

	return sumUpperLeftQuadrant
}

func sum(arr []int32) int32 {
	sum := int32(0)
	for _, num := range arr {
		sum += num
	}
	return sum
}
