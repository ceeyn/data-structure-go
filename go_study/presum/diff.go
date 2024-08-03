package presum

type TwoDiff1 struct {
	diff [][]int
}

// diff[x+1][y+1] 是方便right+1的
func NewTwoDiff1(x, y int) *TwoDiff1 {
	d := &TwoDiff1{
		diff: make([][]int, x+1),
	}
	for i := 0; i < x; i++ {
		d.diff[i] = make([]int, y+1)
	}
	return d
}

// diff[x+2][y+2] 1.是方便right+1的,2.是多加了第一行，第一列，
// 规定diff[0][0]是0，方便求前缀和的
func NewTwoDiff2(x, y int) *TwoDiff1 {
	d := &TwoDiff1{
		diff: make([][]int, x+2),
	}
	for i := 0; i < x; i++ {
		d.diff[i] = make([]int, y+2)
	}
	return d
}

// [left,right] + x
func OneDiff(left int, right int, arr []int, x int) []int {
	var flags []int
	// +1是为了如果right为最后一个元素时方便标记
	flags = make([]int, len(arr)+1)
	flags[left] += x
	flags[right+1] -= x
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += flags[i]
		arr[i] += sum
	}
	return arr
}

// 这里采取这种方式：diff[x+2][y+2]
func TwoDiff(x1, y1, x2, y2 int, arr [][]int, x int) [][]int {
	var flags [][]int
	// +1是为了如果right为最后一个元素时方便标记
	flags = make([][]int, len(arr)+2)
	for i := 0; i < len(arr[0]); i++ {
		flags[i] = make([]int, len(arr[0])+2)
	}
	flags[x1+1][y1+1] += x
	flags[x1+1][y2+2] -= x
	flags[x2+2][y1+1] -= x
	flags[x2+2][y2+2] += x
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			flags[i+1][j+1] += flags[i+1][j] + flags[i][j+1] - flags[i][j]
		}
	}
	for i := 1; i <= len(arr); i++ {
		for j := 1; j <= len(arr[0]); j++ {
			arr[i-1][j-1] = flags[i][j]
		}
	}
	return arr
}
