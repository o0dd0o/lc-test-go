package main

type NumMatrix struct {
	matr [][]int
	sum  [][]int
	cal  bool
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{matr: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.cal {
		sum := this.sum[row2][col2]
		if row1 > 0 {
			sum -= this.sum[row1-1][col2]
		}
		if col1 > 0 {
			sum -= this.sum[row2][col1-1]
		}
		if col1 > 0 && row1 > 0 {
			sum += this.sum[row1-1][col1-1]
		}
		return sum
	}

	lenx := len(this.matr[0])
	leny := len(this.matr)
	sumMatr := [][]int{}
	for y := 0; y < leny; y++ {
		sumMatr = append(sumMatr, make([]int, lenx))
		for x := 0; x < lenx; x++ {
			tm := this.matr[y][x]
			if x > 0 {
				tm += sumMatr[y][x-1]
			}
			if y > 0 {
				tm += sumMatr[y-1][x]
			}
			if x > 0 && y > 0 {
				tm -= sumMatr[y-1][x-1]
			}
			sumMatr[y][x] = tm
		}
	}
	this.sum = sumMatr
	this.cal = true
	return this.SumRegion(row1, col1, row2, col2)
}
func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	println(obj.SumRegion(2, 1, 4, 3))
	println(obj.SumRegion(1, 1, 2, 2))
	println(obj.SumRegion(1, 2, 2, 4))
}
