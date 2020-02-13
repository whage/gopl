package surface

import (
	"fmt"
	"math"
	"io"
	//"os"
)

const (
	width, height = 600, 320 // canvas size in pixels
	cells = 15 // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale = height * 0.3 // pixels per z unit
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Point3D struct {
	X, Y, Z float64
}

type Point2D struct {
	X, Y float64
}

type Quad3D struct {
	A, B, C, D Point3D
}

type Quad2D struct {
	A, B, C, D Point2D
}

func (q *Quad3D) GetAverageHeight() float64 {
	return (q.A.Z + q.B.Z + q.C.Z + q.D.Z) / 4
}

func project3DPoint(x, y, z float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	//x := xyrange * (float64(i)/cells)
	//y := xyrange * (float64(j)/cells)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	//fmt.Fprintln(os.Stderr, sy)

	if math.IsNaN(sy) {
		//sy = 0
	}

	return sx, sy
}

func (q *Quad3D) ProjectTo2D() Quad2D {
	ax, ay := project3DPoint(q.A.X, q.A.Y, q.A.Z)
	bx, by := project3DPoint(q.B.X, q.B.Y, q.B.Z)
	cx, cy := project3DPoint(q.C.X, q.C.Y, q.C.Z)
	dx, dy := project3DPoint(q.D.X, q.D.Y, q.D.Z)

	return Quad2D{
		Point2D{ax, ay},
		Point2D{bx, by},
		Point2D{cx, cy},
		Point2D{dx, dy},
	}
}

func Draw(w io.Writer) {
	var minZ float64 = 0
	var maxZ float64 = 0

	polygons := make([]Quad3D, 0)

	for i := -15.0; i < cells; i++ {
		for j := -15.0; j < cells; j++ {
			q := Quad3D{
				Point3D{i+1, j, f(i+1, j)},
				Point3D{i, j, f(i, j)},
				Point3D{i, j+1, f(i, j+1)},
				Point3D{i+1, j+1, f(i+1, j+1)},
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

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7;' "+
		"width='%d' height='%d'>", width, height)

	for _, q := range polygons {	
		redComponent := getRedComponentAt(q.GetAverageHeight(), minZ, maxZ) 
		blueComponent := getBlueComponentAt(q.GetAverageHeight(), minZ, maxZ)	
		p := q.ProjectTo2D()

		fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%02x00%02x;'/>\n", p.A.X, p.A.Y, p.B.X, p.B.Y, p.C.X, p.C.Y, p.D.X, p.D.Y, redComponent, blueComponent)
	}

	fmt.Fprintf(w, "</svg>")
	//fmt.Fprintf(w, "<p>Z range: %f -> %f</p>", minZ, maxZ)
}

func getRedComponentAt(z, minZ, maxZ float64) int {
	return int((z - minZ) * 255/(maxZ - minZ))
}

func getBlueComponentAt(z, minZ, maxZ float64) int {
	return int(255 - (z - minZ) * 255/(maxZ - minZ))
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
