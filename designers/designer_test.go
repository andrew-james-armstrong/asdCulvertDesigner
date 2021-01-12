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
	r := ReinforcementLayer{0.15, 0.025, 0.032, 0.65}
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
	r := ReinforcementLayer{0.15, 0.025, 0.032, 0.65}
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
	r := ReinforcementLayer{0.15, 0.025, 0.032, 0.65}
	want := "<Layer<LargeBar>0.032</LargeBar><SmallBar>0.025</SmallBar><Spacing>0.15</Spacing><Level>0.65</Level></Layer>"
	have := r.ToXML()
	if have != want {
		t.Fatal(fmt.Sprintf("\nHave %s\nWant %s", have, want))
	}
}

func TestSectionToXML(t *testing.T) {
	r := ReinforcementLayer{0.15, 0.025, 0.032, 0.65}
	s := RectangularConcreteSection{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, r, r, r, r}
	s.Breadth = 1.5
	s.Height = 0.45
	want := fmt.Sprintf("<Rectangular_concrete_section><breadth>%f</Breadth><Height>%f</Height><ConcreteGrade>%f</ConcreteGrade><ReinforcementGrade>%f</ReinforcementGrade><ShearLinkAllowance>%f</ShearLinkAllowance><MinimumCover>%f</MinimumCover><CoverTolerance>%f</CoverTolerance><ReinforcementLayer>%s</ReinforcementLayer><ReinforcementLayer>%s</ReinforcementLayer><ReinforcementLayer>%s</ReinforcementLayer><ReinforcementLayer>%s</ReinforcementLayer></Rectangular_concrete_section>", s.Breadth, s.Height, s.ConcreteGrade, s.ReinforcementGrade, s.ShearLinkAllowance, s.MinCover, s.CoverTolerance, s.ReinfB1.ToXML(), s.ReinfB2.ToXML(), s.ReinfT1.ToXML(), s.ReinfT2.ToXML())
	have := s.ToXML()
	if have != want {
		t.Fatal(fmt.Sprintf("\nHave %s\nWant %s", have, want))
	}
}

func TestMatrixMult(t *testing.T) {
	m1 := creatMatrix(3,2)
	m2 := createMatrix(2,3)
	m1[0] = {3,2}
	m1[1] = {6.15}
	m1[2] = {8,19}
	m2[0] = {13,8,2}
	m2[1] = {6.15,8}
	want :={{51,54,22},{168,273,168},{224,300,330}}
	multiplyMatrices(m1, m2, have)
	if have != want {
		t.Fatal(fmt.Sprintf("\nHave %s\nWant %s", have, want))
	}
}
