package test4

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Test17() {
	cmap := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) //按空格拆分单词
	for in.Scan() {
		if r := in.Text(); r == `quit` {
			break
		} else {
			cmap[r]++
		}
	}
	for k, v := range cmap {
		fmt.Println(k, v)
	}
}

//测试创建一个mao
func Test16() {
	mp := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(mp)
	mp2 := map[string]bool{
		"a": true,
		"b": false,
	}
	if !mp2["b"] {
		fmt.Println(mp2["a"])
	}

	str := []string{"asd", "qwe", "11qwe", "bb22"}
	fmt.Println(str)
	//对slice排序（Ascll）
	sort.Strings(str)
	fmt.Println(str)
}

//将多个空格替换为1个空格
func Test15() {
	str := "你好，     世界！"
	by2 := []byte(str)
	for i := 0; i < len(by2)-1; i++ {
		if uint32(by2[i]) <= unicode.MaxLatin1 {
			if by2[i] == ' ' && by2[i+1] == ' ' {
				copy(by2[i:], by2[i+1:])
				by2 = by2[:len(by2)-1]
				i--
			}
		}
	}
	str2 := string(by2)
	fmt.Println(str2)
	by := []rune(str)
	for i, j := 0, len(by)-1; i < j; i, j = i+1, j-1 {
		by[i], by[j] = by[j], by[i]
	}
	for i := 0; i < len(by)-1; i++ {
		if unicode.IsSpace(by[i]) && unicode.IsSpace(by[i+1]) {
			copy(by[i:], by[i+1:])
			by = by[:len(by)-1]
			i--
		}
	}
	str = string(by)
	fmt.Println(str)

}

//消除相邻的两个元素重复的字符串
func Test14() {
	str := []string{"aa", "bb", "cc", "dd", "dd", "ee", "ff", "ff"}
	str = func() []string {
		for i := 0; i < len(str)-1; i++ {
			if str[i] == str[i+1] {
				str[i], str[i+1] = "", ""
				i++
			}
		}
		return str
	}()
	fmt.Println(str)
}

//数组指针的反转函数,数组指针指向第一个元素的指针
func Test13() {
	s1 := []int{1, 2, 3, 4, 5}
	rotate(s1, 5)
}

func rotate(s []int, i int) {
	for x := 0; x < i; x++ {
		//s[4], s[:4] = s[0], s[1:5]
		s = append(s[1:len(s)], s[0])
	}
	fmt.Println(s)
	// tmp := []int{}
	// tmp = s[:i]
	// s[:len(s)-i] = s[i:]
}

func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//测试函数是否会改变slice的值
func Test12() {
	str := []string{2: "aaa", 5: "bbb", 11: "ccc"}
	fmt.Println(str, len(str))
	fmt.Println(getNotNil2(str))
	fmt.Println(str, len(str))
}

//会覆盖
func getNotNil(str []string) []string {
	i := 0
	for _, s := range str {
		if s != "" {
			str[i] = s
			i++
		}
	}
	return str[:i]
}

//不会
func getNotNil2(str []string) []string {
	str2 := str[:0]
	for _, s := range str {
		if s != "" {
			str2 = append(str2, s)
		}
	}
	return str2
}

func Test11() {
	sz := [3]int{0, 1, 2}
	sc := []int{0, 1, 2, 3}
	fmt.Println(&sz[0], &sc[0])
	fmt.Println(sz[0] == sc[0])
	str := "你好，世界！"
	sc2 := []rune(str)
	fmt.Println(sc2)
	fmt.Printf("%q\n", sc2)
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//测试数组与slice的区别
//用数组部分元素生成的slice默认容量是开始位到数组结束
func Test10() {
	//定义数组
	sz := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	sc := sz[2:5]
	fmt.Println(sz, cap(sz))
	fmt.Println(sc, cap(sc))
	fmt.Printf("sz=%T,sc=%T\n", sz, sc)
	sc = append(sc, 1, 2, 3, 4)
	fmt.Println(sc, cap(sc))
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
		//性能太低，应该有更优的方法
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
