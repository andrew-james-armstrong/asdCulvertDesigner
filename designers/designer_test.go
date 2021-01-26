package designers

import (
	"fmt"
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	nodes := make([]Point, 2)
	nodes[0] = Point{0, 0, 0}
	nodes[1] = Point{3, 4, 0}
	want := 5.0
	have := nodes[0].DistanceBetween(nodes[1])
	if have != want {
		t.Fatal(fmt.Sprintf("Have %f\nwant %f", have, want))
	}
}

func TestPointToXML(t *testing.T) {
	node := Point{3, 4, 0}
	want := "<point n=\"1\"><x>3.000000</x><y>4.000000</y><z>0.000000</z></point>"
	have := node.ToXML(1)
	if have != want {
		t.Fatal(fmt.Sprintf("Have %s\n want %s", have, want))
	}
}

func TestSectionArea(t *testing.T) {
	r := ReinforcementLayer{"T1", 0.15, 0.025, 0.032, 0.65}
	s := RectangularConcreteSection{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, r, r, r, r}
	s.Breadth = 1.5
	s.Height = 0.45
	want := 0.675
	have := s.Area()
	if have != want {
		t.Fatal(fmt.Sprintf("Have %8.7f\nwant %8.7f", have, want))
	}
}

func TestSectionInertia(t *testing.T) {
	r := ReinforcementLayer{"T1", 0.15, 0.025, 0.032, 0.65}
	s := RectangularConcreteSection{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, r, r, r, r}
	s.Breadth = 1.5
	s.Height = 0.45
	want := 0.011390625
	have := s.Inertia()
	diff := math.Abs(have - want)
	if diff > 0.000001 {
		t.Fatal(fmt.Sprintf("\nHave %12.11f\nWant %12.11f", have, want))
	}
}

func TestBarLayerToXML(t *testing.T) {
	r := ReinforcementLayer{"T1", 0.15, 0.025, 0.032, 0.65}
	want := "<Layer><Name>T1</Name><LargeBar>0.032</LargeBar><SmallBar>0.025</SmallBar><Spacing>0.15</Spacing><Level>0.65</Level></Layer>"
	have := r.ToXML()
	if have != want {
		t.Fatal(fmt.Sprintf("\nHave %s\nWant %s", have, want))
	}
}

func TestActiveEarthPressures(t *testing.T) {
	//Test ka coefficient
	want := float64(1 / 3.0)
	have := ka(30)
	if math.Abs(have-want) > 0.000000001 {
		t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have, want))
	}
}
func TestPassiveEarthPressures(t *testing.T) {
	//Test kp coefficient
	want := float64(3.0)
	have := kp(30)
	if math.Abs(have-want) > 0.000000001 {
		t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have, want))
	}
}
func TestRestEarthPressures(t *testing.T) {
	//Test k0 coefficient
	want := 0.5
	have := k0(30)
	if math.Abs(have-want) > 0.000000001 {
		t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have, want))
	}
}

func TestSectionToXML(t *testing.T) {
	r := ReinforcementLayer{"T1", 0.15, 0.025, 0.032, 0.65}
	s := RectangularConcreteSection{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, r, r, r, r}
	s.Breadth = 1.5
	s.Height = 0.45
	want := fmt.Sprintf("<Rectangular_concrete_section><Breadth>%f</Breadth><Height>%f</Height><ConcreteGrade>%f</ConcreteGrade><ReinforcementGrade>%f</ReinforcementGrade><ShearLinkAllowance>%f</ShearLinkAllowance><MinimumCover>%f</MinimumCover><CoverTolerance>%f</CoverTolerance><ReinforcementLayers>%s%s%s%s</ReinforcementLayers></Rectangular_concrete_section>", s.Breadth, s.Height, s.ConcreteGrade, s.ReinforcementGrade, s.ShearLinkAllowance, s.MinCover, s.CoverTolerance, s.ReinfB1.ToXML(), s.ReinfB2.ToXML(), s.ReinfT1.ToXML(), s.ReinfT2.ToXML())
	have := s.ToXML()
	if have != want {
		t.Fatal(fmt.Sprintf("\nHave %s\nWant %s", have, want))
	}
}

func TestMatrixMult(t *testing.T) {
	m1 := [][]float64{{3, 2}, {6, 15}, {8, 9}}
	m2 := [][]float64{{13, 8, 2}, {6, 15, 8}}
	want := [][]float64{{51, 54, 22}, {168, 273, 132}, {158, 199, 88}}
	have, _ := multiplyMatrices(m1, m2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if have[i][j] != want[i][j] {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestIdentityMatrix(t *testing.T) {
	want := createMatrix(3, 3)
	want[0][0] = 1
	want[1][1] = 1
	want[2][2] = 1
	have := createIdentityMatrix(3, 3)
	//	fmt.Println("have:", have)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if have[i][j] != want[i][j] {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

/*
func TestMatrixTransform(t *testing.T) {
	trans := XZRotationMatrix(math.Pi / 4)
	origin := [][]float64{{4, 1, 3, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {3, 4, 5, 0, 0, 0}, {0, 0, 0, 4, 1, 3}, {0, 0, 0, 0, 1, 0}, {0, 0, 0, 3, 4, 5}}
	want := [][]float64{{-3, 1, 4, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {-5, 4, 3, 0, 0, 0}, {0, 0, 0, -3, 1, 4}, {0, 0, 0, 0, 1, 0}, {0, 0, 0, -5, 4, 3}}
	have, _ := multiplyMatrices(origin, trans)
	fmt.Println("origin:", origin)
	fmt.Println("trans:", trans)
	fmt.Println("want:", want)
	fmt.Println("have:", have)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}
*/

func TestFEM_UDL(t *testing.T) {
	udl := 30.0
	length := 5.0
	want := [][]float64{{0}, {udl * length / 2}, {-udl * length * length / 12}, {0}, {udl * length / 2}, {udl * length * length / 12}}
	have := FEM_UDL(udl, length)
	//	fmt.Println("have:", have)
	//	fmt.Println("want:", want)
	for i := 0; i < 6; i++ {
		for j := 0; j < 1; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestFEM_Point(t *testing.T) {
	point := 30.0
	length := 5.0
	start := 1.5
	want := [][]float64{{0}, {point * (length - start) / length}, {-point * start * (length - start) * (length - start) / length / length}, {0}, {point * start / length}, {point * start * start * (length - start) / length / length}}
	have := FEM_Point(point, start, length)
	//	fmt.Println("have:", have)
	//	fmt.Println("want:", want)
	for i := 0; i < 6; i++ {
		for j := 0; j < 1; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestFEM_Linear(t *testing.T) {
	udl := 30.0
	length := 5.0
	want := [][]float64{{0}, {udl * length / 3.0}, {-udl * length * length / 30}, {0}, {udl * length / 6}, {udl * length * length / 20}}
	have := FEM_Linear(udl, length)
	//	fmt.Println("have:", have)
	//	fmt.Println("want:", want)
	for i := 0; i < 6; i++ {
		for j := 0; j < 1; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestFEM_Patch(t *testing.T) {
	udl := 30.0
	length := 5.0
	start := 0.0
	finish := 5.0
	want := FEM_UDL(udl, length)
	have := FEM_Patch(udl, start, finish, length)
	for i := 0; i < 6; i++ {
		for j := 0; j < 1; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestMergeStiffnessMatrices(t *testing.T) {
	mainStiffness := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
	}
	elementStiffness := [][]float64{
		{20, 20, 20, -20, -20, -20},
		{-10, -10, -10, 10, 10, 10},
		{15, 15, 15, -15, -15, -15},
		{21, 21, 21, 22, 22, 22},
		{30, 30, 30, 35, 35, 35},
		{-100, -100, -100, 50, 50, 50},
	}
	want := [][]float64{
		{20, 20, 20, 0, 0, 0, -20, -20, -20},
		{-9, -9, -9, 1, 1, 1, 11, 11, 11},
		{17, 17, 17, 2, 2, 2, -13, -13, -13},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{28, 28, 28, 7, 7, 7, 29, 29, 29},
		{38, 38, 38, 8, 8, 8, 43, 43, 43},
		{-91, -91, -91, 9, 9, 9, 59, 59, 59},
	}
	//	fmt.Println("want:", want)
	have := mergeStiffnessMatrices(mainStiffness, elementStiffness, 0, 2)
	//	fmt.Println("have:", have)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestReverseMergeStiffnessMatrices(t *testing.T) {
	mainStiffness := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
	}
	elementStiffness := [][]float64{
		{20, 20, 20, -20, -20, -20},
		{-10, -10, -10, 10, 10, 10},
		{15, 15, 15, -15, -15, -15},
		{21, 21, 21, 22, 22, 22},
		{30, 30, 30, 35, 35, 35},
		{-100, -100, -100, 50, 50, 50},
	}
	want := [][]float64{
		{22, 22, 22, 0, 0, 0, 21, 21, 21},
		{36, 36, 36, 1, 1, 1, 31, 31, 31},
		{52, 52, 52, 2, 2, 2, -98, -98, -98},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{-13, -13, -13, 7, 7, 7, 27, 27, 27},
		{18, 18, 18, 8, 8, 8, -2, -2, -2},
		{-6, -6, -6, 9, 9, 9, 24, 24, 24},
	}
	//	fmt.Println("want:", want)
	have := mergeStiffnessMatrices(mainStiffness, elementStiffness, 2, 0)
	//	fmt.Println("have:", have)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestElementMatrix(t *testing.T) {
	a := 0.5
	i := 2.0
	e := 10.0
	l := 5.0
	want := [][]float64{{1, 0, 0, -1, 0, 0}, {0, 240 / 125, 120 / 25, 0, -240 / 125, 120 / 25}, {0, 120 / 25, 80 / 5, 0, -120 / 25, 40 / 5}, {-1, 0, 0, 1, 0, 0}, {0, -240 / 125, -120 / 25, 0, 240 / 125, -120 / 25}, {0, 120 / 25, 40 / 5, 0, -120 / 250, -80 / 5}}
	have := createElementMatrix(a, i, e, l)
	for i := 0; i < 6; i++ {
		for j := 0; j < 1; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

/*
func TestEncastBeam(t *testing.T) {
	mat := createMatrix(15, 15)
	force := createMatrix(15, 1)
	a := 0.5
	i := 0.03
	e := 35000.0
	l := 5.0 / 4
	m := createElementMatrix(a, i, e, l)
	f := FEM_UDL(10, l)
	mergeStiffnessMatrices(mat, m, 0, 1)
	mergeStiffnessMatrices(mat, m, 1, 2)
	mergeStiffnessMatrices(mat, m, 2, 3)
	mergeStiffnessMatrices(mat, m, 3, 4)
	mat[0][0] = mat[0][0] * 1000000.0
	mat[12][12] = mat[12][12] * 1000000.0
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			force[3*i+j][0] = force[3*i+j][0] + f[j][0]
		}
	}
	inv, sing := inverseMatrix(mat)
	if sing != nil {
		fmt.Println(sing)
	}
	fmt.Println("mat:", mat)
	fmt.Println("inverse:", inv)
	fmt.Println("force:", force)
}
*/
func TestTranspose(t *testing.T) {
	a := [][]float64{{4.3, 78.2, 23.9}, {56.3, 87.2, 129.0}, {93.1, 55.3, 72.3}}
	want := [][]float64{{4.3, 56.3, 93.1}, {78.2, 87.2, 55.3}, {23.9, 129.0, 72.3}}
	have := Transpose(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}

func TestParabola(t *testing.T) {
	a := [][]float64{{120 * 120, 17, 100}, {130 * 130, 15, 110}, {145 * 145, 12, 95}}
	inv, _ := inverseMatrix(a)
	fmt.Println("Inv: ", inv)
	b := [][]float64{{101}, {103.8}, {98.9}}
	fmt.Println("B: ", b)
	k, _ := multiplyMatrices(inv, b)
	fmt.Println("K: ", k)
	x := [][]float64{{99 * 99}, {99}, {1}}
	y := k[0][0]*x[0][0] + k[1][0]*x[1][0] + k[2][0]*x[2][0]
	fmt.Println("Y:", y)
	want := [][]float64{{100}, {110}, {99}}
	//	fmt.Println("Have:", have)
	fmt.Println("Want:", want)
}

/*
func Test_LDL_Decpompose(t *testing.T) {
	a := createMatrix(3, 3)
	a = [][]float64{{4.3, 78.2, 23.9}, {78.2, 87.2, 129.0}, {23.9, 129.0, 72.3}}
	L, D := LDLDecompose(a)
	want, _ := inverseMatrix(a)
	have, _ := multiplyMatrices(L, D)
	have, _ = multiplyMatrices(have, Transpose(L))
	fmt.Println("Lower:", L)
	fmt.Println("Diagonal:", D)
	fmt.Println("Transpose:", Transpose(L))
	fmt.Println("want:", want)
	fmt.Println("have:", have)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.0000000001 {
				t.Fatal(fmt.Sprintf("\nHave %f\nWant %f", have[i][j], want[i][j]))
			}
		}
	}
}
*/

func Test_LLtDecpompose(t *testing.T) {
	S := createMatrix(15, 15)
	a := createElementMatrix(0.5, 0.003, 35000.0, 2.5)
	fmt.Println(a[0][0])
	for i := 0; i < 4; i++ {
		mergeStiffnessMatrices(S, a, i, i+1)
	}
	fmt.Println(S[3][3])
	S[0][0] = S[0][0] * 10000
	S[1][1] = S[1][1] * 10000
	S[2][2] = S[2][2] * 10000
	S[12][12] = S[12][12] * 10000
	S[14][14] = S[14][14] * 10000
	L, T := LLtDecompose(S)
	for i := 0; i < len(S); i++ {
		fmt.Print(L[3][i], " ")
	}
	fmt.Print("\n")
	for i := 0; i < len(S); i++ {
		fmt.Print(T[3][i], " ")
	}
	fmt.Print("\n")
	fmt.Println(L[3][3])
	fmt.Println(T[3][3])
	want := S
	have, _ := multiplyMatrices(L, T)
	fmt.Println("want:", want)
	fmt.Println("have:", have)
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(S); j++ {
			if math.Abs(have[i][j]-want[i][j]) > 0.00001 {
				t.Fatal(fmt.Sprintf("\nIndex[%v][%v]\nHave %f\nWant %f", i, j, have[i][j], want[i][j]))
			}
		}
	}
}
