package designers

import (
	"github.com/andrew-james-armstrong/asdCulvertDesigner/utils"
	"log"
	"math"
	"net/http"
)

func two_cell(req *http.Request) string {
	var structure Culvert
	var section CulvertSection
	var defaultReinforcementLayer = ReinforcementLayer{"T1", 0.15, 0.0, 0.0, 0.0}
	var structureBackfill GroundProperties

	project := "<project_name>" + req.FormValue("Pname") + "</project_name><structure_name>" + req.FormValue("Sname") + "</structure_name><structure_reference>" + req.FormValue("Sref") + "</structure_reference><structure_number>" + req.FormValue("SNo") + "</structure_number>"

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

	// finer details at corners
	section.TopHaunchWidth = extract(req, "top_haunch_width")
	section.TopHaunchHeight = extract(req, "top_haunch_height")
	section.BottomHaunchWidth = extract(req, "bottom_haunch_width")
	section.BottomHaunchHeight = extract(req, "bottom_haunch_height")

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

	fill_cover_to_roof := structure.IPGroundLevel - section.Height - structure.IPInvertLevel - section.RoofThickness
	if fill_cover_to_roof < 0.6 {
		log.Fatal("Inadequate cover to fit structure.") // No point going any further if the structure wont fit
	}
	structure.LeftInvertLevel = structure.IPInvertLevel + structure.LeftLength*structure.LongitudinalGrade
	structure.RightInvertLevel = structure.IPInvertLevel - structure.RightLength*structure.LongitudinalGrade
	//	left_founding_level := structure.LeftInvertLevel - section.BaseThickness
	//	right_founding_level := structure.RightInvertLevel - section.BaseThickness

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
	structureBackfill.KA = (1 - math.Sin(structureBackfill.InternalFrictionAngle)) / (1 + math.Cos(structureBackfill.InternalFrictionAngle))
	structureBackfill.KP = (1 + math.Sin(structureBackfill.InternalFrictionAngle)) / (1 - math.Cos(structureBackfill.InternalFrictionAngle))
	structureBackfill.K0 = (1 - math.Sin(structureBackfill.InternalFrictionAngle)) / (1 + math.Sin(structureBackfill.InternalFrictionAngle))

	// Loading details
	Fill_udl = fill_cover_to_roof * structureBackfill.BulkDensity

	cover_to_top := IP_Ground_level - IP_Invert_level - Height - Rth
	downstream_invert_level := IP_Invert_level - downstream_length*structure_longitudinal_grade
	upstream_invert_level := IP_Invert_level + upstream_length*structure_longitudinal_grade

	// Build node list
	nodes := section.GenerateTwoCellNodes()

	// Build Element list
	elements := make([]Element, len(nodes))

	elements[0] = Element{0, 1, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[1] = Element{1, 2, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[2] = Element{2, 3, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[3] = Element{3, 4, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[4] = Element{4, 5, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[5] = Element{5, 6, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[6] = Element{6, 7, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[7] = Element{7, 8, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[8] = Element{8, 9, RectangularConcreteSection{section.Length, section.RoofThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[9] = Element{9, 10, RectangularConcreteSection{section.Length, section.RoofThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[10] = Element{10, 11, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[11] = Element{11, 12, RectangularConcreteSection{section.Length, section.RoofThickness + section.TopHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[12] = Element{12, 13, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[13] = Element{13, 14, RectangularConcreteSection{section.Length, section.WallThickness + section.TopHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[14] = Element{14, 15, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[15] = Element{15, 16, RectangularConcreteSection{section.Length, section.WallThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[16] = Element{16, 17, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[17] = Element{17, 18, RectangularConcreteSection{section.Length, section.WallThickness + section.BottomHaunchWidth, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[18] = Element{18, 19, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[19] = Element{19, 20, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[20] = Element{20, 21, RectangularConcreteSection{section.Length, section.BaseThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[21] = Element{21, 22, RectangularConcreteSection{section.Length, section.BaseThickness, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[22] = Element{22, 23, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight/2, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}
	elements[23] = Element{23, 0, RectangularConcreteSection{section.Length, section.BaseThickness + section.BottomHaunchHeight, Fcu, Fy, 0.01, 0.03, 0.005, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer, defaultReinforcementLayer}}

	designResult := make([]utils.Parameter, 12, 50)
	designResult[0] = utils.Parameter{"Cover to top of culvert", cover_to_top, 0, utils.LENGTH, "Dim"}
	designResult[1] = utils.Parameter{"Downstream invert level", downstream_invert_level, 0, utils.LENGTH, "Dim"}
	designResult[2] = utils.Parameter{"Upstream invert level", upstream_invert_level, 0, utils.LENGTH, "Dim"}
	designResult[3] = utils.Parameter{"Wall cross-sectional area", wall_section_area, 0, utils.SQUARE, "Area"}
	designResult[4] = utils.Parameter{"Roof cross-sectional area", roof_section_area, 0, utils.SQUARE, "Area"}
	designResult[5] = utils.Parameter{"Base cross-sectional area", base_section_area, 0, utils.SQUARE, "Area"}
	designResult[6] = utils.Parameter{"Whole culvert section area", culvert_section_area, 0, utils.SQUARE, "Area"}
	designResult[7] = utils.Parameter{"Whole culvert perimeter", culvert_section_perimeter, 0, utils.LENGTH, "Length"}
	designResult[8] = utils.Parameter{"Wall section 2nd moment of area", wall_section_inertia, 0, utils.INERTIA, "Inertia"}
	designResult[9] = utils.Parameter{"Roof section 2nd moment of area", roof_section_inertia, 0, utils.INERTIA, "Inertia"}
	designResult[10] = utils.Parameter{"Base section 2nd moment of area", base_section_inertia, 0, utils.INERTIA, "Inertia"}

	nodelist := "<Node>"
	for i := 0; i < len(nodes); i++ {
		nodelist = nodelist + nodes[i].ToXML(i)
	}
	nodelist = nodelist + "</Node>"
	elementlist := ""
	for i := 0; i < len(elements); i++ {
		elementlist = elementlist + elements[i].ToXML(i)
	}
	resultlist := ""
	for i := 0; i < len(designResult); i++ {
		resultlist = resultlist + designResult[i].ToXML()
	}
	xml_head := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><?xml-stylesheet type=\"text/xsl\" href=\"./assets/calculationReport.xsl\"?>"
	return xml_head + "<Calculation><Project>" + project + "</Project>" + nodelist + elementlist + "<Results>" + resultlist + "</Results></Calculation>"
}
