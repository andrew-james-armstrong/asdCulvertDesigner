package designers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const DESIGNER_VERSION = "v0.0.1"

var projectName, StructureName, structureRef, structureNumber string
var IP_Ground_level, IP_Invert_level, structure_longitudinal_grade, structure_skew, upstream_length, downstream_length float64
var bed_thickness, water_depth, freeboard, Analysis_length, Top_haunch_width, Top_haunch_height, Bottom_haunch_width, Bottom_haunch_height float64
var Width, Height, Wth, Rth, Bth, Lth, Th, Bh, Fcu, Ec, Es, Fy, Gamma_c, Gamma_s float64
var wall_section_area, roof_section_area, base_section_area, culvert_seection_area, culvert_section_perimeter, wall_section_inertia, roof_section_inertia float64
var base_section_inertia, culvert_section_area, culvert_sectio_perimeter float64
var Fill_udl, Superimposed_dead_udl, Vehicle_udl, Abnormal_vehicle_udl, Permanent_load_factor, Live_load_factor, Abnormal_load_factor float64
var Servicability_limit_state_factor, Ultimate_limit_state_factor, Earth_pressure_factor_max, Earth_pressure_factor_min float64
var Title string

type Point struct {
	X float64
	Y float64
	Z float64
}

func (self Point) Add(p Point) Point {
	return Point{
		self.X + p.X,
		self.Y + p.Y,
		self.Z + p.Z,
	}
}

func (self Point) Subtract(p Point) Point {
	return Point{
		self.X - p.X,
		self.Y - p.Y,
		self.Z - p.Z,
	}
}

func (self Point) ScaleBy(s float64) Point {
	return Point{
		self.X * s,
		self.Y * s,
		self.Z * s,
	}
}

func (self Point) DistanceBetween(p Point) float64 {
	sum := (p.X-self.X)*(p.X-self.X) +
		(p.Y-self.Y)*(p.Y-self.Y) +
		(p.Z-self.Z)*(p.Z-self.Z)

	return math.Sqrt(sum)
}

func (self Point) ToXML(n int) string {
	return fmt.Sprintf("<point n=\"%d\"><x>%f</x><y>%f</y><z>%f</z></point>", n, self.X, self.Y, self.Z)
}

type ReinforcementLayer struct {
	BarSpacing  float64
	SmallBarDia float64
	LargeBarDia float64
	LayerCentre float64
}

func (self ReinforcementLayer) NextBarSizeIncrementUpwards() ReinforcementLayer {
	if self.SmallBarDia < self.LargeBarDia {
		self.SmallBarDia = self.LargeBarDia
	} else {
		var bars = []float64{0, 0.006, 0.01, 0.012, 0.016, 0.020, 0.025, 0.032, 0.040, 0.05}
		for i, v := range bars {
			if self.LargeBarDia == v {
				self.LargeBarDia = bars[i+1]
			}
		}
	}
	return self
}

func (self ReinforcementLayer) ToXML() string {
	return fmt.Sprintf("<Layer<LargeBar>%v</LargeBar><SmallBar>%v</SmallBar><Spacing>%v</Spacing><Level>%v</Level></Layer>", self.LargeBarDia, self.SmallBarDia, self.BarSpacing, self.LayerCentre)
}

type RectangularConcreteSection struct {
	Breadth            float64
	Height             float64
	ConcreteGrade      float64
	ReinforcementGrade float64
	ShearLinkAllowance float64
	MinCover           float64
	CoverTolerance     float64
	ReinfB1            ReinforcementLayer
	ReinfB2            ReinforcementLayer
	ReinfT1            ReinforcementLayer
	ReinfT2            ReinforcementLayer
}

func (self RectangularConcreteSection) Area() float64 {
	return self.Breadth * self.Height
}

func (self RectangularConcreteSection) Inertia() float64 {
	return (self.Breadth * self.Height * self.Height * self.Height) / 12.0
}

func (self RectangularConcreteSection) ToXML() string {
	return fmt.Sprintf("<Rectangular_concrete_section><breadth>%f</Breadth><Height>%f</Height><ConcreteGrade>%f</ConcreteGrade><ReinforcementGrade>%f</ReinforcementGrade><ShearLinkAllowance>%f</ShearLinkAllowance><MinimumCover>%f</MinimumCover><CoverTolerance>%f</CoverTolerance><ReinforcementLayer>%s</ReinforcementLayer><ReinforcementLayer>%s</ReinforcementLayer><ReinforcementLayer>%s</ReinforcementLayer><ReinforcementLayer>%s</ReinforcementLayer></Rectangular_concrete_section>", self.Breadth, self.Height, self.ConcreteGrade, self.ReinforcementGrade, self.ShearLinkAllowance, self.MinCover, self.CoverTolerance, self.ReinfB1.ToXML(), self.ReinfB2.ToXML(), self.ReinfT1.ToXML(), self.ReinfT2.ToXML())
}

type Culvert struct {
	IPGroundLevel     float64
	IPInvertLevel     float64
	LongitudinalGrade float64
	SkewAngle         float64
	LeftLength        float64
	RightLength       float64
	LeftInvertLevel   float64
	RightInvertLevel  float64
}

type CulvertSection struct {
	Width              float64
	WallThickness      float64
	RoofThickness      float64
	BaseThickness      float64
	Freeboard          float64
	WaterDepth         float64
	BedDepth           float64
	Length             float64
	Height             float64
	TopHaunchWidth     float64
	TopHaunchHeight    float64
	BottomHaunchWidth  float64
	BottomHaunchHeight float64
}

func (self CulvertSection) GenerateOneCellNodes() []Point {
	// Build node list
	// Note this is a closed section so node 24 is the same point as node 0
	nodes := make([]Point, 23, 23)

	// First set up some convenient intermediate values
	mid_wall := (self.Height - self.BaseThickness - self.RoofThickness - self.TopHaunchHeight - self.BottomHaunchHeight) / 2
	mid_roof := self.Width/2 - self.TopHaunchWidth
	mid_base := self.Width/2 - self.BottomHaunchWidth

	// Do a walk around the perimeter to generate each node

	nodes[0] = Point{0, 0, 0}
	nodes[1] = nodes[0].Add(Point{0, self.BaseThickness / 2, 0})
	nodes[2] = nodes[1].Add(Point{0, self.BottomHaunchHeight, 0})
	nodes[3] = nodes[2].Add(Point{0, mid_wall, 0})
	nodes[4] = nodes[3].Add(Point{0, mid_wall, 0})
	nodes[5] = nodes[4].Add(Point{0, self.TopHaunchHeight, 0})
	nodes[6] = nodes[5].Add(Point{0, self.RoofThickness / 2, 0})
	nodes[7] = nodes[6].Add(Point{self.WallThickness / 2, 0, 0})
	nodes[8] = nodes[7].Add(Point{self.TopHaunchWidth, 0, 0})
	nodes[9] = nodes[8].Add(Point{mid_roof, 0, 0})
	nodes[10] = nodes[9].Add(Point{mid_roof, 0, 0})
	nodes[11] = nodes[10].Add(Point{self.TopHaunchWidth, 0, 0})
	nodes[12] = nodes[11].Add(Point{self.WallThickness / 2, 0, 0})
	nodes[13] = nodes[13].Subtract(Point{0, self.RoofThickness / 2, 0})
	nodes[14] = nodes[14].Subtract(Point{0, self.TopHaunchHeight, 0})
	nodes[15] = nodes[14].Subtract(Point{0, mid_wall, 0})
	nodes[16] = nodes[15].Subtract(Point{0, mid_wall, 0})
	nodes[17] = nodes[16].Subtract(Point{0, self.BottomHaunchHeight, 0})
	nodes[18] = nodes[17].Subtract(Point{0, self.BaseThickness / 2, 0})
	nodes[19] = nodes[18].Subtract(Point{self.WallThickness / 2, 0, 0})
	nodes[20] = nodes[19].Subtract(Point{self.BottomHaunchWidth, 0, 0})
	nodes[21] = nodes[20].Subtract(Point{mid_base, 0, 0})
	nodes[22] = nodes[21].Subtract(Point{mid_base, 0, 0})
	nodes[23] = nodes[22].Subtract(Point{self.BottomHaunchWidth, 0, 0})
	return nodes
}

func (self CulvertSection) GenerateTwoCellNodes() []Point {
	// Build node list
	// Note this is a closed section so node 24 is the same point as node 0
	nodes := make([]Point, 23, 23)

	// First set up some convenient intermediate values
	mid_wall := (self.Height - self.BaseThickness - self.RoofThickness - self.TopHaunchHeight - self.BottomHaunchHeight) / 2
	mid_roof := self.Width/2 - self.TopHaunchWidth
	mid_base := self.Width/2 - self.BottomHaunchWidth

	// Do a walk around the perimeter to generate each node

	nodes[0] = Point{0, 0, 0}
	nodes[1] = nodes[0].Add(Point{0, self.BaseThickness / 2, 0})
	nodes[2] = nodes[1].Add(Point{0, self.BottomHaunchHeight, 0})
	nodes[3] = nodes[2].Add(Point{0, mid_wall, 0})
	nodes[4] = nodes[3].Add(Point{0, mid_wall, 0})
	nodes[5] = nodes[4].Add(Point{0, self.TopHaunchHeight, 0})
	nodes[6] = nodes[5].Add(Point{0, self.RoofThickness / 2, 0})
	nodes[7] = nodes[6].Add(Point{self.WallThickness / 2, 0, 0})
	nodes[8] = nodes[7].Add(Point{self.TopHaunchWidth, 0, 0})
	nodes[9] = nodes[8].Add(Point{mid_roof, 0, 0})
	nodes[10] = nodes[9].Add(Point{mid_roof, 0, 0})
	nodes[11] = nodes[10].Add(Point{self.TopHaunchWidth, 0, 0})
	nodes[12] = nodes[11].Add(Point{self.WallThickness / 2, 0, 0})
	nodes[13] = nodes[13].Subtract(Point{0, self.RoofThickness / 2, 0})
	nodes[14] = nodes[14].Subtract(Point{0, self.TopHaunchHeight, 0})
	nodes[15] = nodes[14].Subtract(Point{0, mid_wall, 0})
	nodes[16] = nodes[15].Subtract(Point{0, mid_wall, 0})
	nodes[17] = nodes[16].Subtract(Point{0, self.BottomHaunchHeight, 0})
	nodes[18] = nodes[17].Subtract(Point{0, self.BaseThickness / 2, 0})
	nodes[19] = nodes[18].Subtract(Point{self.WallThickness / 2, 0, 0})
	nodes[20] = nodes[19].Subtract(Point{self.BottomHaunchWidth, 0, 0})
	nodes[21] = nodes[20].Subtract(Point{mid_base, 0, 0})
	nodes[22] = nodes[21].Subtract(Point{mid_base, 0, 0})
	nodes[23] = nodes[22].Subtract(Point{self.BottomHaunchWidth, 0, 0})
	return nodes
}

func extract(req *http.Request, s string) float64 {
	val, err := strconv.ParseFloat(req.FormValue(s), 64)
	if err != nil {
		println(s)
		log.Fatal(err)
	}
	return val
}
