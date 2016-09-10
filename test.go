package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	Test1()
}

//测试，编写一个字符串大小写转换工具
//测试编写一个输入/opt/qwe/test.go转换为test，com.qwe.test.go转换为com.qwe.test
//了解strings包的相关功能

//测试字符串区字节值
func Test1() {
	str := "世22"
	//fmt.Println(str[0:6] + str[0:6])
	str2 := "\u4e1622"
	fmt.Println(str == str2, str2) //true
	str3 := "世22界A"
	fmt.Println(str3[:len(str)] == str) //true
	str4 := "22"
	fmt.Println(strings.Contains(str3, str4)) //包含 true

	fmt.Printf("查看str3的utf8长度【%d】和unicode长度【%d】\n", len(str3), utf8.RuneCountInString(str3))
	//按uncicode的方式取得字符，以防止汉字被拆分报错
	for i := 0; i < len(str3); {
		r, size := utf8.DecodeRuneInString(str3[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	n := 0
	//简洁的方式
	for i, s := range str3 {
		//fmt.Printf("%q\n", s)
		fmt.Printf("%d\t%q\t%d\n", i, s, s)
		n++
	}
	//fmt.Println(n)
	r, size := utf8.DecodeRuneInString(str3)
	fmt.Println(r, size, str3)      //19990 3
	fmt.Println(unicode.ToLower(r)) //19990

	//测试字符串与byte切片互转
	s1 := "abc"
	b1 := []byte(s1)
	s2 := string(b1)
	fmt.Println(s1, b1, s2)
	//计算字符串2在字符串1中存在的数量
	fmt.Println(strings.Count(str3, "23"))
	//转换为数组按空格
	l := strings.Fields("2 324,322 33 4")
	fmt.Println(len(l))
	//擦看字符串2在字符串1中的位置
	fmt.Println(strings.Index(str3, "22"))
	//按指定的间隔符号连接数组
	fmt.Println(strings.Join(l, "sep"))

	i2 := [3]int{1, 2, 3}
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range i2 {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	fmt.Printf("%s\n", buf.String())
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
