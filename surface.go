package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 // canvas size in pixels
	cells = 50 // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale = height * 0.3 // pixels per z unit
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Point struct {
	X, Y, Z float64
}

type Quad struct {
	A, B, C, D Point
}

func (q *Quad) GetAverageHeight() float64 {
	return (q.A.Z + q.B.Z + q.C.Z + q.D.Z) / 4
}

func main() {
	var minZ float64 = 0
	var maxZ float64 = 0

	polygons := make([]Quad, 0)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			q := Quad{
				Point{ax, ay, az},
				Point{bx, by, bz},
				Point{cx, cy, cz},
				Point{dx, dy, dz},
			}

			polygons = append(polygons, q)

			averageZ := q.GetAverageHeight()

			if averageZ < minZ {
				minZ = averageZ
			}

			if averageZ > maxZ {
				maxZ = averageZ
			}
		}
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)

	for _, p := range polygons {
		redComponent := getRedComponentAt(p.GetAverageHeight(), minZ, maxZ) 
		blueComponent := getBlueComponentAt(p.GetAverageHeight(), minZ, maxZ)
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%02x00%02x';/>\n", p.A.X, p.A.Y, p.B.X, p.B.Y, p.C.X, p.C.Y, p.D.X, p.D.Y, redComponent, blueComponent)
	}

	fmt.Println("</svg>")
	fmt.Printf("<p>Z range: %f -> %f</p>", minZ, maxZ)
}

func getRedComponentAt(z, minZ, maxZ float64) int {
	return int((z - minZ) * 255/(maxZ - minZ))
}

func getBlueComponentAt(z, minZ, maxZ float64) int {
	return int(255 - (z - minZ) * 255/(maxZ - minZ))
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onti 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
	//return math.Pow(x, 2)/100 - math.Pow(y, 2)/100
	//return math.Pow(x, 3)/100 - 3*x*math.Pow(y, 2)/100
}
