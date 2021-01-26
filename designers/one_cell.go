package designers

import (
	"fmt"
	"github.com/andrew-james-armstrong/asdCulvertDesigner/utils"
	_ "log"
	_ "math"
	"net/http"
	"time"
)

const CHARGERATE = 1000000

type Element struct {
	Start, End int
	Section    RectangularConcreteSection
}

func (self Element) ToXML(n int) string {
	return fmt.Sprintf("<Element n=\"%d\"><start>%d</start><end>%d</end>%s</Element>", n, self.Start, self.End, self.Section.ToXML())
}

type GroundProperties struct {
	BulkDensity           float64
	InternalFrictionAngle float64
	WallFrictionAngle     float64
	TopOfLayerLevel       float64
	GroundwaterLevel      float64
	KA                    float64
	K0                    float64
	KP                    float64
}

func (self GroundProperties) ToXML() string {
	return fmt.Sprintf("<Soil><BulkDensity>%v</BulkDensity><InternalFrictionAngle>%v</InternalFrictionAngle><WallFrictionAngle>%v</WallFrictionAngle><TopOfLayerLevel>%v</TopOfLayerLevel><GroundwaterLevel>%v</GroundwaterLevel><KA>%v</KA><KP>%v</KP><K0>%v</K0></Soil>", self.BulkDensity, self.InternalFrictionAngle, self.WallFrictionAngle, self.TopOfLayerLevel, self.GroundwaterLevel, self.KA, self.KP, self.K0)
}

type ElementLoad struct {
	LoadcaseNo                 int
	ElementNo                  int
	StartPosition, EndPosition int // This is the distance from end1 that the load starts or ends
	StartLoad, EndLoad         float64
	LoadType                   string // This is UDL,Point, Patch
}

type Loadcase struct {
	Name        string
	Number      int
	ElementList []int
	Factor      float64
}

type LoadCombination struct {
	Name      string
	Factor    float64
	Loadcases []int
}

type LoadEnvelope struct {
	Name         string
	Factor       float64
	Combinations []int
}

func initialiseLayers(def ReinforcementLayer) []ReinforcementLayer {
	layer := make([]ReinforcementLayer, 4)
	layer[0] = def
	layer[1] = def
	layer[2] = def
	layer[3] = def
	layer[0].Name = "T1"
	layer[1].Name = "T2"
	layer[2].Name = "B1"
	layer[3].Name = "B2"
	layer[0].LayerCentre = -0.1
	layer[1].LayerCentre = -0.05
	layer[2].LayerCentre = 0.1
	layer[3].LayerCentre = 0.15
	return layer
}

func One_Cell(req *http.Request) string {
	var structure Culvert
	var section CulvertSection
	var defaultReinforcementLayer = ReinforcementLayer{"T1", 0.15, 0.006, 0.010, 0.100}
	var structureBackfill GroundProperties
	var client string
	var account_no string

	client = "The Disney Corporation"
	account_no = "1234-0345-6279-1287"
	project := "<project_name>" + req.FormValue("Pname") + "</project_name><structure_name>" + req.FormValue("Sname") + "</structure_name><structure_reference>" + req.FormValue("Sref") + "</structure_reference><structure_number>" + req.FormValue("SNo") + "</structure_number>"

	start_time := time.Now()

	//basic setting out of structure
	structure = Culvert{
		extract(req, "ground_level"),
		extract(req, "invert_level"),
		extract(req, "structure_longitudinal_grade"),
		extract(req, "structure_skew"),
		extract(req, "upstream_length"),
		extract(req, "downstream_length"),
		0, //To be added later
		0, //To be added later
		0, //To be added later
	}

	//Structure cross section details
	section = CulvertSection{
		extract(req, "width"),
		extract(req, "wall_thickness"),
		extract(req, "roof_thickness"),
		extract(req, "base_thickness"),
		extract(req, "freeboard"),
		extract(req, "water_depth"),
		extract(req, "bed_thickness"),
		extract(req, "length"),
		0, //To be added below
		0, //To be added later
		0, //To be added later
		0, //To be added later
		0, //To be added later
	}
	// Internal height of structure calculated from the parameters supplied
	section.Height = section.BedDepth + section.WaterDepth + section.Freeboard
	// finer details at corners
	section.TopHaunchWidth = extract(req, "top_haunch_width")
	section.TopHaunchHeight = extract(req, "top_haunch_height")
	section.BottomHaunchWidth = extract(req, "bottom_haunch_width")
	section.BottomHaunchHeight = extract(req, "bottom_haunch_height")

	structure.RoofCover = structure.IPGroundLevel - section.Height - structure.IPInvertLevel - section.RoofThickness
	if structure.RoofCover < 0.6 {
		println(fmt.Sprintf("%f cover is insufficient to fit structure and pavement.", structure.RoofCover)) // No point going any further if the structure wont fit
	} else {
		println(fmt.Sprintf("Fill cover depth above roof slab is %8.3f.", structure.RoofCover))
	}
	structure.LeftInvertLevel = structure.IPInvertLevel + structure.LeftLength*structure.LongitudinalGrade
	structure.RightInvertLevel = structure.IPInvertLevel - structure.RightLength*structure.LongitudinalGrade
	left_founding_level := structure.LeftInvertLevel - section.BaseThickness
	right_founding_level := structure.RightInvertLevel - section.BaseThickness

	//structure materials
	Fcu = extract(req, "fcu")
	Gamma_c = extract(req, "gamma_c")
	Gamma_s = extract(req, "gamma_s")
	Es = extract(req, "es")
	Ec = extract(req, "ec")
	Fy = extract(req, "fy")

	//Ground condition information
	structureBackfill = GroundProperties{
		extract(req, "gamma_fill"),
		extract(req, "theta"),
		extract(req, "delta"),
		0,
		0,
		0,
		0,
		0,
	}
	structureBackfill.KA = ka(structureBackfill.InternalFrictionAngle)
	structureBackfill.KP = kp(structureBackfill.InternalFrictionAngle)
	structureBackfill.K0 = k0(structureBackfill.InternalFrictionAngle)

	println(fmt.Sprintf("Active EP = %8.3f", structureBackfill.KA))
	println(fmt.Sprintf("Passive EP = %8.3f", structureBackfill.KP))
	println(fmt.Sprintf("At-Rest EP = %8.3f", structureBackfill.K0))
	// Loading details
	Fill_udl = structure.RoofCover * structureBackfill.BulkDensity
	superimposed_dead_udl := extract(req, "w_superimposed")
	vehicle_udl := extract(req, "vehicle_udl")
	abnormal_vehicle_udl := extract(req, "abnormal_udl")

	println(fmt.Sprintf("Super dead %8.3f", superimposed_dead_udl))
	println(fmt.Sprintf("Vehicle UDL %8.3f", vehicle_udl))
	println(fmt.Sprintf("Abnormal UDL %8.3f", abnormal_vehicle_udl))

	loadcases := make([]Loadcase, 50)
	loadcases[0] = Loadcase{"Self Weight", 1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, 1.15}
	loadcases[1] = Loadcase{"Superimposed load surfacing + fill", 2, []int{6, 7, 8, 9, 10, 11}, 1.75}
	loadcases[2] = Loadcase{"Min earth pressure left", 3, []int{0, 1, 2, 3, 4, 5}, 1.2}
	loadcases[3] = Loadcase{"At-rest earth pressure left", 4, []int{0, 1, 2, 3, 4, 5}, 1.2}
	loadcases[4] = Loadcase{"Max earth pressure left", 5, []int{0, 1, 2, 3, 4, 5}, 1.2}
	loadcases[5] = Loadcase{"Min earth pressure right", 6, []int{12, 13, 14, 15, 16, 17}, 1.2}
	loadcases[6] = Loadcase{"At-rest earth pressure right", 7, []int{12, 13, 14, 15, 16, 17}, 1.2}
	loadcases[7] = Loadcase{"Max earth pressure right", 8, []int{12, 13, 14, 15, 16, 17}, 1.2}
	loadcases[8] = Loadcase{"Live load 1 left", 9, []int{6}, 1.5}
	loadcases[9] = Loadcase{"Live load 2 left", 10, []int{6, 7}, 1.5}
	loadcases[10] = Loadcase{"Live load 3 left", 11, []int{6, 7, 8}, 1.5}
	loadcases[11] = Loadcase{"Live load 4 left", 12, []int{6, 7, 8, 9}, 1.5}
	loadcases[12] = Loadcase{"Live load 5 left", 13, []int{6, 7, 8, 9, 10}, 1.5}
	loadcases[13] = Loadcase{"Live load 6 left", 14, []int{6, 7, 8, 9, 10, 11}, 1.5}
	loadcases[14] = Loadcase{"Live load 7 right", 15, []int{7, 8, 9, 10, 11}, 1.5}
	loadcases[15] = Loadcase{"Live load 8 right", 16, []int{8, 9, 10, 11}, 1.5}
	loadcases[16] = Loadcase{"Live load 9 right", 17, []int{9, 10, 11}, 1.5}
	loadcases[17] = Loadcase{"Live load 10 right", 18, []int{10, 11}, 1.5}
	loadcases[18] = Loadcase{"Live load 11 right", 19, []int{11}, 1.5}
	loadcases[19] = Loadcase{"Abnormal load 1 left", 20, []int{6}, 1.3}
	loadcases[20] = Loadcase{"Abnormal load 2 left", 21, []int{6, 7}, 1.3}
	loadcases[21] = Loadcase{"Abnormal load 3 left", 22, []int{6, 7, 8}, 1.3}
	loadcases[22] = Loadcase{"Abnormal load 4 left", 23, []int{6, 7, 8, 9}, 1.3}
	loadcases[23] = Loadcase{"Abnormal load 5 left", 24, []int{6, 7, 8, 9, 10}, 1.3}
	loadcases[24] = Loadcase{"Abnormal load 6 left", 25, []int{6, 7, 8, 9, 10, 11}, 1.3}
	loadcases[25] = Loadcase{"Abnormal load 7 right", 26, []int{7, 8, 9, 10, 11}, 1.3}
	loadcases[26] = Loadcase{"Abnormal load 8 right", 27, []int{8, 9, 10, 11}, 1.3}
	loadcases[27] = Loadcase{"Abnormal load 9 right", 28, []int{9, 10, 11}, 1.3}
	loadcases[28] = Loadcase{"Abnormal load 10 right", 29, []int{10, 11}, 1.3}
	loadcases[29] = Loadcase{"Abnormal load 11 right", 30, []int{11}, 1.3}
	loadcases[30] = Loadcase{"Internal water pressure hydrostatic", 31, []int{1, 2, 3, 4, 5, 12, 13, 14, 15, 16}, 1.1}
	loadcases[31] = Loadcase{"Internal water pressure max", 32, []int{19, 20, 21, 22}, 1.1}
	loadcases[32] = Loadcase{"Internal temperature rise", 33, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, 0.8}
	loadcases[33] = Loadcase{"Internal temperature fall", 34, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, 0.8}

	// Need to look up the Eurocode equation 6.10 etc in EN 1990
	/// Need to incorporate the psi factors for leading and associated effects.
	combinations := make([]LoadCombination, 10)
	combinations[0] = LoadCombination{"SW + Max EP Left+ Min EP right + Live", 1.0, []int{0, 1, 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	combinations[1] = LoadCombination{"SW + Min EP Left+ Max EP right + Live", 1.0, []int{0, 1, 2, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	combinations[2] = LoadCombination{"SW + Max EP Left+ Min EP right + Abnormal", 1.0, []int{0, 1, 4, 5, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}}
	combinations[3] = LoadCombination{"SW + Min EP Left+ Max EP right + Abnormal", 1.0, []int{0, 1, 2, 7, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}}
	combinations[4] = LoadCombination{"SW + At-Rest EP + Live + WP", 1.0, []int{0, 1, 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 30, 31}}
	combinations[5] = LoadCombination{"SW + At-Rest EP + Live", 1.0, []int{0, 1, 2, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	combinations[6] = LoadCombination{"SW + At-Rest EP + Abnormal + WP", 1.0, []int{0, 1, 4, 5, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}}
	combinations[7] = LoadCombination{"SW + At-Rest EP + Abnormal", 1.0, []int{0, 1, 2, 7, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}}
	combinations[8] = LoadCombination{"SW + At-Rest EP + Live + TempRise", 1.0, []int{0, 1, 2, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 32}}
	combinations[9] = LoadCombination{"SW + At-Rest EP + Live + TempFall", 1.0, []int{0, 1, 2, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 33}}

	envelopes := make([]LoadEnvelope, 10)
	envelopes[0] = LoadEnvelope{"Unfactored", 1.0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	envelopes[1] = LoadEnvelope{"SLS", 1.05, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	envelopes[2] = LoadEnvelope{"ULS", 1.15, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}

	// Calculate horizontal earth perssures at the levels of the top and bottom frame elements
	/*
		top_active_earth_pressure := structureBackfill.KA * (fill_cover_to_roof + Rth/2) * structureBackfill.BulkDensity
		bottom_active_earth_pressure := structureBackfill.KA * (fill_cover_to_roof + Height + Rth + Bth/2) * structureBackfill.BulkDensity
		top_passive_earth_pressure := structureBackfill.KP * (fill_cover_to_roof + Rth/2) * structureBackfill.BulkDensity
		bottom_passive_earth_pressure := structureBackfill.KP * (fill_cover_to_roof + Height + Rth + Bth/2) * structureBackfill.BulkDensity
		top_at_rest_earth_pressure := structureBackfill.K0 * (fill_cover_to_roof + Rth/2) * structureBackfill.BulkDensity
		bottom_at_rest_earth_pressure := structureBackfill.K0 * (fill_cover_to_roof + Height + Rth + Bth/2) * structureBackfill.BulkDensity
	*/

	//Design factors of safety
	Permanent_load_factor = extract(req, "dead")
	Earth_pressure_factor_max = extract(req, "earth_max")
	Earth_pressure_factor_min = extract(req, "earth_min")
	Live_load_factor = extract(req, "live")
	Abnormal_load_factor = extract(req, "live2")
	Servicability_limit_state_factor = extract(req, "SLS")
	Ultimate_limit_state_factor = extract(req, "ULS")

	nodes := section.GenerateOneCellNodes()
	// Build Element list
	elements := make([]Element, len(nodes))
	layer := initialiseLayers(defaultReinforcementLayer)

	elements[0] = Element{0, 1, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[1] = Element{1, 2, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[2] = Element{2, 3, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[3] = Element{3, 4, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[4] = Element{4, 5, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[5] = Element{5, 6, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[6] = Element{6, 7, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[7] = Element{7, 8, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[8] = Element{8, 9, RectangularConcreteSection{section.Length, section.RoofThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[9] = Element{9, 10, RectangularConcreteSection{section.Length, section.RoofThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[10] = Element{10, 11, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[11] = Element{11, 12, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[12] = Element{12, 13, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[13] = Element{13, 14, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[14] = Element{14, 15, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[15] = Element{15, 16, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[16] = Element{16, 17, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[17] = Element{17, 18, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[18] = Element{18, 19, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[19] = Element{19, 20, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[20] = Element{20, 21, RectangularConcreteSection{section.Length, section.BaseThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[21] = Element{21, 22, RectangularConcreteSection{section.Length, section.BaseThickness, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[22] = Element{22, 23, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}
	elements[23] = Element{23, 0, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, layer[0], layer[1], layer[2], layer[3]}}

	designResult := make([]utils.Parameter, 50)
	designResult[0] = utils.Parameter{"Cover to top of culvert = Ground Level - Invert Level - structure internal height - structure roof thickness.", structure.RoofCover, 0, utils.LENGTH, "Dim"}
	designResult[1] = utils.Parameter{"Left headwall invert level = Invert at IP + structure gradient * structure length at left hand side.", structure.LeftInvertLevel, 0, utils.LENGTH, "Dim"}
	designResult[2] = utils.Parameter{"Left headwall founding level = Left headwall invert level - base slab thickness.", left_founding_level, 0, utils.LENGTH, "Dim"}
	designResult[3] = utils.Parameter{"Right headwall invert level = Invert at IP - structure gradient * structure length at right hand side.", structure.RightInvertLevel, 0, utils.LENGTH, "Dim"}
	designResult[4] = utils.Parameter{"Right headwall founding level = Right headwall inver level - base slab thickness", right_founding_level, 0, utils.LENGTH, "Dim"}
	designResult[5] = utils.Parameter{"Wall cross-sectional area", wall_section_area, 0, utils.SQUARE, "Area"}
	designResult[6] = utils.Parameter{"Roof cross-sectional area", roof_section_area, 0, utils.SQUARE, "Area"}
	designResult[7] = utils.Parameter{"Base cross-sectional area", base_section_area, 0, utils.SQUARE, "Area"}
	designResult[8] = utils.Parameter{"Whole culvert section area", culvert_section_area, 0, utils.SQUARE, "Area"}
	designResult[9] = utils.Parameter{"Whole culvert perimeter", culvert_section_perimeter, 0, utils.LENGTH, "Length"}
	designResult[10] = utils.Parameter{"Wall section 2nd moment of area", wall_section_inertia, 0, utils.INERTIA, "Inertia"}
	designResult[11] = utils.Parameter{"Roof section 2nd moment of area", roof_section_inertia, 0, utils.INERTIA, "Inertia"}
	designResult[12] = utils.Parameter{"Base section 2nd moment of area", base_section_inertia, 0, utils.INERTIA, "Inertia"}

	nodelist := "<Nodes>"
	for i := 0; i < len(nodes); i++ {
		nodelist = nodelist + nodes[i].ToXML(i)
	}
	nodelist = nodelist + "</Nodes>"

	elementlist := "<Elements>"
	for i := 0; i < len(elements); i++ {
		elementlist = elementlist + elements[i].ToXML(i)
	}
	elementlist = elementlist + "</Elements>"
	resultlist := ""
	for i := 0; i < len(designResult); i++ {
		resultlist = resultlist + designResult[i].ToXML()
	}

	finish_time := time.Now()
	cost := CHARGERATE * (finish_time.Sub(start_time)).Seconds()
	println(cost)

	charge := fmt.Sprintf("<Charge Client=\"%s\" Account=\"%s\"ItemCost=\"%6.3f\"/>", client, account_no, cost)

	xml_head := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><?xml-stylesheet type=\"text/xsl\" href=\"./assets/calculationReport.xsl\"?>"
	metadata := "<Version>" + DESIGNER_VERSION + "</Version><Date>" + DESIGNER_VERSION_DATE + "</Date><Name>" + DESIGNER_NAME + "</Name><RunTime>" + DESIGNER_RUN_TIME + "</RunTime>"
	return xml_head + "<Calculation><Meta>" + metadata + "</Meta><Project>" + project + "</Project>" + nodelist + elementlist + "<Results>" + resultlist + "</Results>" + charge + "</Calculation>"
}
