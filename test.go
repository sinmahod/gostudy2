package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	Test9()
}

//练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
func Test9() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(c1)
	fmt.Println(c2)
	var it int = func() int {
		mp := make(map[int]byte)
		for i, cs2 := range c2 {
			mp[i] = cs2
		}
		var i int
		for _, cs1 := range c1 {
			for k, v := range mp {
				if cs1 == v {
					fmt.Printf("存在重复的为%d\n", v)
					i++
					delete(mp, k)
					break
				}
			}
		}
		return i
	}()
	fmt.Printf("%d", it)
}

//测试...声明数组：根据初始化的值数量确定数组大小
func Test8() {
	i := [...]int{1, 2, 3, 4}
	fmt.Println(i, len(i))
	s := [...]int{99: 1}
	fmt.Println(s, len(s))
	ptr := [...]byte{19: 1}
	fmt.Println(ptr)
	updateByte(&ptr) //如果直接传入指针的话会修改变量内容
	fmt.Println(ptr)

}

func updateByte(ptrs *[20]byte) {
	*ptrs = [20]byte{}
}

//练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func Test7() {
	stra := "abcssas"
	strb := "cbaasss"
	if len(stra) != len(strb) {
		fmt.Printf("%s与%s字符长度不一致\n", stra, strb)
		return
	}

	mp := make(map[int]rune)

	for _, s1 := range stra {
		isexists := false
		for i, s2 := range strb {
			_, ok := mp[i] //判断这个元素是否已经使用过了
			if s1 == s2 && !ok {
				mp[i] = s2 //相同的元素放到map里
				isexists = true
				break
			}
		}
		if !isexists {
			fmt.Printf("%s与%s字符不一致\n", stra, strb)
			return
		}
	}
	fmt.Printf("%s与%s字符一致\n", stra, strb)
}

func Test6() {
	st := comma("1234567")
	fmt.Println(st)
}

func comma(s string) string {
	var buf bytes.Buffer
	con := len(s)
	i := len(s)
	for idx := 0; idx < len(s); idx++ {
		i--
		str := s[idx]
		fmt.Println(str)
		fmt.Fprintf(&buf, "%s", string(str))
		if (idx+1)%3 == 0 && idx+1 < con {
			buf.WriteString(",")
		}

	}
	return buf.String()
	// n := len(s)
	// if n <= 3 {
	// 	return s
	// }
	// return comma(s[:n-3]) + "," + s[n-3:]
}

//了解strings包的相关功能
func Test1() {
	//测试，字符串大小写转换工具
	str := "hello,world!"
	str = strings.ToUpper(str)
	fmt.Println(str)
	//测试编写一个输入/opt/qwe/test.go转换为test，com.qwe.test.go转换为com.qwe.test
	fmt.Println(getFileName("com.qwe.test.go"))
	fmt.Println(filepath.Split("/opt/qwe/test.go"))
}

func getFileName(str string) string {
	idx := strings.LastIndex(str, "/")
	idx2 := strings.LastIndex(str, ".")
	str2 := str
	if idx != -1 {
		str2 = str2[idx+1:]
	}
	if idx2 != -1 {
		str2 = str2[:idx2]
	}
	return str2
}

//测试字符串区字节值
func Test2() {
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
