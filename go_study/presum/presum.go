package presum

type PreSum struct {
	sum [][]int
}

func NewPreSum(vals [][]int) *PreSum {
	m := len(vals)
	if m == 0 {
		return &PreSum{sum: [][]int{}}
	}
	n := len(vals[0])
	sum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = vals[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}
	return &PreSum{sum}
}

func (sum *PreSum) GetSum(x1 int, y1 int, x2 int, y2 int) int {
	return sum.sum[x2+1][y2+1] - sum.sum[x2+1][y1] - sum.sum[x1][y2+1] + sum.sum[x1][y1]
}
