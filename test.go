package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	Test5()
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

//测试gif动画生成  $ go run test.go > test.gif
func Test5() {
	//获取19位时间戳  ,纳秒数
	//将纳秒数当作种子，种子的作用就好象一个key，使用同一个种子得到的随机数永远是一样的，计算速度根据服务器配置有关系，计算机1纳秒15次左右
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//使用ReadFile函数读取文件全部内容
func Test4() {
	files := os.Args[1:]
	if len(files) > 0 {
		for _, f := range files {
			b, err := ioutil.ReadFile(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "发现错误：%v\n", err)
				continue
			}
			fmt.Println("文件的内容为:")
			fmt.Printf("%s\n", b)
			line := strings.Split(string(b), "\n")
			fmt.Printf("文件共有%d行\n", len(line))
		}
	} else {
		fmt.Println("命令行语句中没有发现文件")
	}
}

//读取命令行中的文件并打开将文件内容按行输出   使用Open要记得Close文件
func Test3() {
	mp := make(map[int]string)
	if len(os.Args) == 1 {
		readFile(os.Stdin, mp)
	} else {
		for i := 1; i < len(os.Args); i++ {
			f, err := os.Open(os.Args[i])
			if err != nil {
				fmt.Println(err)
				continue
			}
			readFile(f, mp)
			f.Close()
		}
	}

	for i := 1; i <= len(mp); i++ {
		fmt.Println(i, mp[i])
	}
}

func readFile(f *os.File, mp map[int]string) {
	input := bufio.NewScanner(f)
	i := 0
	for input.Scan() {
		if input.Err() != nil {
			fmt.Println(input.Err())
			continue
		}
		if input.Text() == `quit` {
			break
		}
		i++
		mp[i] = input.Text()
	}
}

//测试：统计用户输入次数
func Test2() {
	mp := make(map[string]int)
	//以行的形式读取用户输入   os.Stdin ： 命令行输入
	input := bufio.NewScanner(os.Stdin)
	//循环获取用户在命令行中输入的值，直到用户输入quit则退出并判断用户输入的重复次数及重复字符
	for input.Scan() { //开始读取，等待用户输入一行
		if input.Text() == `quit` {
			break
		}
		if input.Err() != nil {
			fmt.Println("输入错误", input.Err())
			return
		} else {
			//如果存在这个key的话则value+1
			mp[input.Text()]++
		}
	}
	for key, value := range mp {
		fmt.Printf("您输入了%s  %d 次 \n", key, value)
	}
}

//测试从命令行读取数据并相加
func Test1() {
	var j int = 0
	//os.Args的第一个元素os.Args[0]为命令行
	for i := 1; i < len(os.Args); i++ {
		//判断是否可以转换为int类型
		c, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println("错误不可以转换为int", os.Args[i])
		} else {
			j += c
		}
	}
	fmt.Printf("计算和，结果为:%d\n", j)
	fmt.Println(join(os.Args[1:]))
}

//测试字符串连接
func join(str []string) string {
	return strings.Join(str, ",")
}
