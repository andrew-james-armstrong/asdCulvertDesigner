package utils

// Set up the structures that hold the specific data types

var OneCellData = PageStruct{
	"Single Cell Box Culvert Design",

	[]Title{
		{"Project Name", "Pname", "Wibble"},
		{"Structure Name", "Sname", "Wobble"},
		{"Structure Reference", "Sref", "des6723876641102a"},
		{"Structure Number", "SNo", "C443786543199874A"},
	},
	[]Group{
		{"Proposed ground level at the intersection point with the main alignment.", "ground_level", 100.0, LENGTH},
		{"Invert level at the intersection point with the main alignment.", "invert_level", 95.0, LENGTH},
		{"Longitudinal gradient required for the structure at the intersection point (Assumes that the flow is from left to right looking up-chainage on the main alignment).", "structure_longitudinal_grade", 0.01, PCT},
		{"Skew angle of culvert at the intersection point with the main alignment. Zero skew = a right angled crossing.", "structure_skew", 0.0, DEG},
		{"Upstream length of culvert from intersection point to start of headwall structure.", "upstream_length", 10.0, LENGTH},
		{"Downstream length of culvert from intersection point to start of headwall structure.", "downstream_length", 10.0, LENGTH},
		{"Culvert internal width face to face of walls.", "width", 5.5, LENGTH},
		{"Required depth of stream bed material.", "bed_thickness", 0.50, LENGTH},
		{"Max flood water depth.", "water_depth", 3.0, LENGTH},
		{"Minimum freeboard height.", "freeboard", 0.6, LENGTH},
		{"Structural wall thickness", "wall_thickness", 0.50, LENGTH},
		{"Structural roof thickness", "roof_thickness", 0.65, LENGTH},
		{"Structural base thickness", "base_thickness", 0.75, LENGTH},
		{"Top haunch width", "top_haunch_width", 0.25, LENGTH},
		{"Top Haunch height", "top_haunch_height", 0.25, LENGTH},
		{"Bottom haunch width", "bottom_haunch_width", 0.25, LENGTH},
		{"Bottom haunch height", "bottom_haunch_height", 0.25, LENGTH},
		{"Analysis Length (defaults to a 1m strip)", "length", 1.00, LENGTH},
	},
	[]Group{
		{"Concrete - Characteristic Strength", "fcu", 35, STRESS},
		{"Concrete - Density", "gamma_c", 2450, DENSITY},
		{"Concrete - Young's Modulus", "ec", 34000, STRESS},
		{"Reinforcement - Characteristic Strength", "fy", 500, STRESS},
		{"Reinforcement - Density", "gamma_s", 7850, DENSITY},
		{"Reinforcement - Young's Modulus", "es", 200000, STRESS},
	},
	[]Group{
		{"Backfill bulk density", "gamma_fill", 18, DENSITY},
		{"Natural ground bulk density", "gamma_ground", 19, DENSITY},
		{"Backfill - Angle of internal friction &#0110;", "theta", 35, DEG},
		{"Backfill - Angle of wall friction delta", "delta", 25, DEG},
		{"Foundation - Allowable bearing pressure", "bearing_pressure", 100, STRESS},
		{"Foundation - Modulus of subgrade reaction", "subgrade_stiffness", 1000, FORCE_LENGTH},
		{"Foundation - Allowable settlement", "max_settlement", 0.01, LENGTH},
	},
	[]Group{
		{"Superimposed UDL", "w_superimposed", 3.5, FORCE_LENGTH},
		{"Vehicle UDL", "vehicle_udl", 10, FORCE_LENGTH},
		{"Abnormal UDL", "abnormal_udl", 20, FORCE_LENGTH},
	},
	[]Group{
		{"Permanent Loads", "dead", 1.15, DIMENSIONLESS},
		{"Earth Pressures", "earth_max", 1.5, DIMENSIONLESS},
		{"Earth Pressures", "earth_min", 0.9, DIMENSIONLESS},
		{"Live Loads", "live", 1.5, DIMENSIONLESS},
		{"Abnormal UDL", "live2", 1.5, DIMENSIONLESS},
		{"Servicability", "SLS", 1.0, DIMENSIONLESS},
		{"Ultimate", "ULS", 1.15, DIMENSIONLESS},
	},
}
