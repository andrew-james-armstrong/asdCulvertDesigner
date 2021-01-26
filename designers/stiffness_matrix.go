package designers

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type symmetricMatrix struct {
	matrixSize, size int
	store            []float64
}

func (self symmetricMatrix) New(matSize int64) {
	size := (matSize - 1) * (matSize)
	//matrixSize := matSize * matSize
	store := make([]float64, size)
	for i := 0; i < len(store); i++ {
		store[i] = 0.0
	}
}

func (self symmetricMatrix) Read(row int64, col int64) float64 {
	if col > row {
		return self.store[col*row]
	} else {
		return self.store[row*col]
	}
}

func (self symmetricMatrix) Write(val float64, row int64, col int64) {
	if col > row {
		self.store[col*row] = val
	} else {
		self.store[row*col] = val
	}
}

func (self symmetricMatrix) PrintStr() string {
	s := "["
	for row := 0; row < self.matrixSize; row++ {
		s += "["
		for col := 0; col < self.matrixSize; col++ {
			if col > row {
				s += fmt.Sprintf("%v,", self.store[col*row])
			} else {
				s += fmt.Sprintf("%v,", self.store[row*col])
			}
			s += "],\n["
		}
	}
	s += "]"
	return s
}

func random(min, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}

func XZRotationMatrix(theta float64) [][]float64 {
	t := createMatrix(6, 6)
	c := math.Cos(theta)
	s := math.Sin(theta)
	t[0][0] = c
	t[0][2] = s
	t[2][0] = -s
	t[2][2] = c
	t[1][1] = 1
	t[3][3] = c
	t[3][5] = s
	t[5][3] = -s
	t[5][5] = c
	t[4][4] = 1
	return t
}

func Transpose(A [][]float64) [][]float64 {
	t := createMatrix(len(A), len(A))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			t[i][j] = A[j][i]
		}
	}
	return t
}

func LLtDecompose(A [][]float64) ([][]float64, [][]float64) {
	L := createMatrix(len(A), len(A))
	for row := 0; row < len(A); row++ { // from row 0 to row len(rows)
		for col := 0; col < len(A); col++ { // From column 0 to column len(cols)
			d := 0.0
			for subcol := 1; subcol < col; subcol++ {
				d = d + L[col][subcol]*L[col][subcol]
			}
			L[col][col] = math.Sqrt(A[col][col] - d)
			if row > col {
				d = 0.0
				for subcol := 1; subcol < col; subcol++ {
					d = d + L[row][subcol]*L[col][subcol]
				}
				L[row][col] = (A[row][col] - d) / L[col][col]
			}
		}
	}
	return L, Transpose(L)
}

func LDLDecompose(A [][]float64) ([][]float64, [][]float64, [][]float64) {
	L := A
	D := createMatrix(len(A), len(A))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			d := 0.0
			for k := 1; k < j; k++ {
				d = d + L[j][k]*L[j][k]*D[k][k]
			}
			D[j][j] = A[j][j] - d
			L[j][j] = 1.0
			if i > j {
				e := 0.0
				for k := 1; k < j; k++ {
					e = e + L[i][k]*L[j][k]*D[k][k]
				}
				L[i][j] = (A[i][j] - e) / D[j][j]
			}
			if j > i {
				L[i][j] = 0.0
			}
		}
	}
	return L, D, Transpose(L)
}

func getCofactor(A [][]float64, temp [][]float64, p int, q int, n int) {
	i := 0
	j := 0

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != q {
				temp[i][j] = A[row][col]
				j++
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func determinant(A [][]float64, n int) float64 {
	D := float64(0)
	if n == 1 {
		return A[0][0]
	}

	temp := createMatrix(n, n)
	sign := 1

	for f := 0; f < n; f++ {
		getCofactor(A, temp, 0, f, n)
		D += float64(sign) * A[0][f] * determinant(temp, n-1)
		sign = -sign
	}

	return D
}

func adjoint(A [][]float64) ([][]float64, error) {
	N := len(A)
	adj := createMatrix(N, N)
	if N == 1 {
		adj[0][0] = 1
		return adj, nil
	}
	sign := 1
	var temp = createMatrix(N, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			getCofactor(A, temp, i, j, N)
			if (i+j)%2 == 0 {
				sign = 1
			} else {
				sign = -1
			}
			adj[j][i] = float64(sign) * (determinant(temp, N-1))
		}
	}
	return adj, nil
}

func inverseMatrix(A [][]float64) ([][]float64, error) {
	N := len(A)
	var inverse = createMatrix(N, N)
	det := determinant(A, N)
	if det == 0 {
		fmt.Println("Singular matrix, cannot find its inverse!")
		return nil, nil
	}

	adj, err := adjoint(A)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			inverse[i][j] = float64(adj[i][j]) / float64(det)
		}
	}

	return inverse, nil
}

func multiplyMatrices(m1 [][]float64, m2 [][]float64) ([][]float64, error) {
	if len(m1[0]) != len(m2) {
		return nil, errors.New("Cannot multiply the given matrices!")
	}

	result := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		result[i] = make([]float64, len(m2[0]))
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return result, nil
}

func createMatrix(row, col int) [][]float64 {
	r := make([][]float64, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i] = append(r[i], 0.0)
		}
	}
	return r
}

func mergeStiffnessMatrices(m1 [][]float64, m2 [][]float64, node1, node2 int) [][]float64 {
	if len(m1) >= node2+3 {
		end1Offset := 3 * node1
		end2Offset := 3 * node2
		for row := 0; row < 3; row++ {
			// Top left quadrant of m2
			for col := 0; col < 3; col++ {
				m1[end1Offset+row][end1Offset+col] = m1[end1Offset+row][end1Offset+col] + m2[row][col]
			}
			//Top right quadrant of m2
			for col := 0; col < 3; col++ {
				m1[end1Offset+row][end2Offset+col] = m1[end1Offset+row][end2Offset+col] + m2[row][col+3]
			}
		}
		for row := 0; row < 3; row++ {
			//Bottom left quadrant of m2
			for col := 0; col < 3; col++ {
				m1[end2Offset+row][end1Offset+col] = m1[end2Offset+row][end1Offset+col] + m2[row+3][col]
			}
			//Bottom right quadrant of m2
			for col := 0; col < 3; col++ {
				m1[end2Offset+row][end2Offset+col] = m1[end2Offset+row][end2Offset+col] + m2[row+3][col+3]
			}
		}
	}
	return m1
}

func createIdentityMatrix(row, col int) [][]float64 {
	r := make([][]float64, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == j {
				r[i] = append(r[i], 1.0)
			} else {
				r[i] = append(r[i], 0.0)
			}
		}
	}
	return r
}

func createElementMatrix(a float64, i float64, e float64, l float64) [][]float64 {
	/*
		Should be
		[ ea/l, 0,        0,  -ea/l,     0,         0]
		[0, 12EI/l^3, 6EI/l^2,  0,   -12EI/l^3, 6EI/l^2]
		[0,  6EI/l^3,   4EI/l,  0,   -6EI/l^2,  2EI/l]
		[-ea/l, 0,        0,    ea/l,     0,         0]
		[0, -12EI/l^3, -6EI/l^2,  0,   12EI/l^3, -6EI/l^2]
		[0,  6EI/l^3,   2EI/l,    0,   -6EI/l^2,  4EI/l]
	*/
	m := createMatrix(6, 6)
	axial := a * e / l
	twoEI := 2.0 * e * i / l
	fourEI := 4.0 * e * i / l
	sixEI := 6.0 * e * i / l / l
	twelveEI := 12.0 * e * i / l / l / l
	m[0][0] = axial
	m[3][3] = axial
	m[0][3] = -axial
	m[3][0] = -axial
	m[1][1] = twelveEI
	m[4][4] = twelveEI
	m[1][4] = -twelveEI
	m[4][1] = -twelveEI
	m[1][2] = sixEI
	m[2][1] = sixEI
	m[4][5] = -sixEI
	m[5][4] = -sixEI
	m[5][1] = sixEI
	m[1][5] = sixEI
	m[2][4] = -sixEI
	m[4][2] = -sixEI
	m[2][2] = fourEI
	m[5][5] = fourEI
	m[2][5] = twoEI
	m[5][2] = twoEI
	return m
}
