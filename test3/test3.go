package test3

import (
	"fmt"
	"math"
	"os"
)

func Test1() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		fmt.Printf("%c\n", x)
	}
}

//得到当前工作空间目录
func Test2() {
	cwd, err := os.Getwd()
	fmt.Println(cwd, err)
}

//测试浮点数超出最大值的反应,超过最大值的时候就会产生误差，导致计算不准确
func Test3() {
	var f float32 = 16777216
	fmt.Println(f == f+1) //true
}

//测试%e
func Test4() {
	fmt.Printf("%e\n", 1.2345400000023) //1.234540e+00
	fmt.Printf("%e\n", 1.23454000000000000200023)
	fmt.Println(math.Exp(1))
	nan := math.NaN()
	fmt.Println(nan == nan)
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Test5() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
