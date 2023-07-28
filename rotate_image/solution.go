package rotate_image

func rotate(matrix [][]int) {
	start := 0
	length := len(matrix) - 1
	for start <= length {
		cycle(matrix, start)
		start++
		length--
	}
}

func cycle(matrix [][]int, start int) {
	length := len(matrix) - start
	for i := start; i < length-1; i++ {
		h, l := start, i
		tmp := matrix[h][l]
		for j := 0; j < 4; j++ {
			newH := l
			newL := len(matrix) - 1 - h

			matrix[newH][newL], tmp = tmp, matrix[newH][newL]
			h, l = newH, newL
		}
	}
}
