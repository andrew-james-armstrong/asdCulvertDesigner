package designers

import (
	_ "fmt"
	_ "math"
)

func FEM_UDL(udl float64, length float64) [][]float64 {
	// [Fx] .. EndA
	// [Fy] .. EndA
	// [Mz] .. EndA
	// [Fx] .. EndB
	// [Fy] .. EndB
	// [Mz] .. EndB
	force := createMatrix(6, 1)
	force[1][0] = udl * length / 2.0
	force[2][0] = -udl * length * length / 12.0
	force[4][0] = udl * length / 2.0
	force[5][0] = udl * length * length / 12.0
	return force
}

func FEM_Point(point float64, start, length float64) [][]float64 {
	// [Fx] .. EndA
	// [Fy] .. End
	// [Mz] .. EndA
	// [Fx] .. EndB
	// [Fy] .. EndB
	// [Mz] .. EndB
	force := createMatrix(6, 1)
	force[1][0] = point * (length - start) / length
	force[4][0] = point * start / length
	force[2][0] = -point * start * (length - start) * (length - start) / length / length
	force[5][0] = point * start * start * (length - start) / length / length
	return force
}

func FEM_Linear(peak_udl float64, length float64) [][]float64 {
	// [Fx] .. EndA
	// [Fy] .. EndA
	// [Mz] .. EndA
	// [Fx] .. EndB
	// [Fy] .. EndB
	// [Mz] .. EndB
	force := createMatrix(6, 1)
	force[1][0] = peak_udl * length / 3
	force[4][0] = peak_udl * length / 6
	force[2][0] = -peak_udl * length * length / 30
	force[5][0] = peak_udl * length * length / 20
	return force
}

func FEM_Patch(udl float64, start, finish, length float64) [][]float64 {
	// [Fx] .. EndA
	// [Fy] .. EndA
	// [Mz] .. EndA
	// [Fx] .. EndB
	// [Fy] .. EndB
	// [Mz] .. EndB
	force := createMatrix(6, 1)
	a := start
	b := finish - start
	c := length - finish
	d := finish
	e := length - start
	force[2][0] = -udl / 12.0 / length / length * (e*e*e*(4*length-3*e) - c*c*c*(4*length-3*c))
	force[5][0] = udl / 12.0 / length / length * (d*d*d*(4*length-3*d) - a*a*a*(4*length-3*a))
	force[1][0] = udl*b*(c+b/2)/length + (force[2][0]+force[5][0])/length
	force[4][0] = udl*b*(a+b/2)/length + (force[5][0]+force[2][0])/length
	return force
}
