package utils

import (
	"fmt"
)

type SI_Unit struct {
	Text   string
	factor float64
}

type Imp_Unit struct {
	Text   string
	factor float64
}

type Parameter struct {
	Description string
	Value       float64
	Default     float64
	Unit        SI_Unit
	Ref         string
}

type Calculation_narrative struct {
	Section_Id string
	Text       string
	Ref        string
}

type CulvertGeomData struct {
	Width   Parameter
	Height  Parameter
	Wth     Parameter
	Rth     Parameter
	Bth     Parameter
	Lth     Parameter
	Fcu     Parameter
	Gamma_c Parameter
	Ec      Parameter
	Fy      Parameter
	Gamma_s Parameter
	Es      Parameter
}

// Set up structures to manage types of input used

type Project struct {
	ProjectName        string
	StructureName      string
	StructureReference string
}

type Group struct {
	Name  string
	Field string
	Value float64
	Unit  SI_Unit
}

type Title struct {
	Name  string
	Field string
	Value string
}

type PageStruct struct {
	Title        string
	Project      []Title
	Geom         []Group
	Property     []Group
	SoilProperty []Group
	Loads        []Group
	Factors      []Group
}

type DesignCalculationReport struct {
	wall_section_area         Parameter
	roof_section_area         Parameter
	base_section_area         Parameter
	culvert_seection_area     Parameter
	culvert_section_perimeter Parameter
	wall_section_inertia      Parameter
	roof_section_inertia      Parameter
	base_section_inertia      Parameter
}

var LENGTH = SI_Unit{"m", 1.0}
var PCT = SI_Unit{"%", 0.01}
var SQUARE = SI_Unit{"m^2", 1.0}
var VOLUME = SI_Unit{"m^3", 1.0}
var INERTIA = SI_Unit{"m^4", 1.0}
var MM = SI_Unit{"mm", 0.001}
var DEG = SI_Unit{"deg", 1.0}
var STRESS = SI_Unit{"N/sq nm", 1.0 - 6}
var DENSITY = SI_Unit{"kg/cu m", 1.0}
var FORCE_LENGTH = SI_Unit{"kN/m", 1000.0}
var DIMENSIONLESS = SI_Unit{"", 1.0}

func (self Group) ToXML(i int) string {
	return fmt.Sprintf("<Record field=\"%s\" Value=\"%f\" Unit=\"%s\">%s</Record>", self.Field, self.Value, self.Unit.Text, self.Name)
}

func (self Parameter) ToXML() string {
	return fmt.Sprintf(`<Parameter value="%f" ref="%s" unit="%s" default="%f">%s</Parameter>`, self.Value, self.Ref, self.Unit.Text, self.Default, self.Description)
}

func getResultList(r []Parameter) string {
	var s string
	s = ""
	for i := range r {
		s += fmt.Sprintf(`<Record Ref="%v" Value="%v" Unit="%v">%v</Record>`, r[i].Ref, r[i].Value, r[i].Unit, r[i].Description)
	}
	return (s)
}
