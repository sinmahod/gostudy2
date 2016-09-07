package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	Test9()
}

//测试web服务器：监听8080端口，用户访问/count则执行counter函数，访问/image则执行getimage，否则执行handlers函数
func Test9() {
	http.HandleFunc("/", handlers)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/image", getimage)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var mp = make(map[string]int)

//为了防止并发导致计数错误我们需要一个同步锁
var mu sync.Mutex

func getimage(w http.ResponseWriter, r *http.Request) {
	var cyc float64
	//解析URL参数，如果不解析r.Form里面是没有值的
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cyc" && len(v) > 0 {
			i, err := strconv.ParseFloat(v[0], 32)
			if err != nil {
				fmt.Fprintf(w, "错误：%s不可以转换为int", v)
				return
			}
			if i > 20 {
				fmt.Fprintf(w, "错误：请修改您的参数。cyc不可以大于20")
				return
			}
			cyc = i
		}
	}
	lissajous(w, cyc)
}

func handlers(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	mu.Lock()
	mp[path]++
	mu.Unlock()
	for k, v := range r.Header {
		//读取参数
		fmt.Fprintln(w, k, v)
	}
	fmt.Fprintf(w, "当前url被访问了%d次\n", mp[path])
	fmt.Fprintln(w, r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "访问地址的Host是:%s，RemoteAddr是%s\n", r.Host, r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		//读取参数
		fmt.Fprintln(w, k, v)
	}

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "当前访问的次数统计为：\n")
	for k, v := range mp {
		fmt.Fprintf(w, "URL：%s 被访问了%d次\n", k, v)
	}
	mu.Unlock()

}

//测试go协程用法：获取多个url的请求时间
func Test8() {
	start := time.Now()
	//创建一个string类型的通道，只有通道类型chan才可以在协程之间通信
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		bl := strings.HasPrefix(url, `http://`)
		if !bl {
			url = `http://` + url
		}
		go rurl(url, ch)
	}
	//用法等同于 _,_ := range os.Args[1:]
	for range os.Args[1:] {
		//代码执行到这里会判断是否有goroutine对通道做send或receive操作，有的话代码就会阻塞在这里，
		//等待goroutine的操作
		fmt.Println(<-ch)
	}
	end := time.Since(start).Seconds()
	fmt.Printf("程序运行总时长%f秒", end)
}

func rurl(url string, ch chan<- string) {
	//得到当前时间
	start := time.Now()
	txt, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("错误的URL：%s,请求访问不到:%s", url, err)
		return
	}
	//ioutil.Discard 可以把这个当成一个垃圾桶，不需要输出的话就用这个。
	by, err2 := io.Copy(ioutil.Discard, txt.Body)
	txt.Body.Close()
	if err2 != nil {
		ch <- fmt.Sprintf("错误的URL：%s,无法读取:%s", url, err)
		return
	}
	ti := time.Since(start).Seconds()
	ch <- fmt.Sprintf("URL：%s 访问用时：%f 秒，字节数%d", url, ti, by)
}

//改用io.Copy方式获取
func Test7() {
	for _, urlstr := range os.Args[1:] {
		//如果没有http://自动不全
		bl := strings.HasPrefix(urlstr, `http://`)
		if !bl {
			urlstr = `http://` + urlstr
		}
		txt, err := http.Get(urlstr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error：%v\n", err)
			//终止进程，并且返回错误代码1
			os.Exit(1)
		}
		_, err2 := io.Copy(os.Stdout, txt.Body)
		//打印状态码
		fmt.Printf("txt.Status:%s\n", txt.Status)
		//防止资源泄漏得到数据后需要关闭
		txt.Body.Close()
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "URL：%s ,Error : %v\n", urlstr, err2)
			//终止进程，并且返回错误代码1
			os.Exit(1)
		}
	}
}

//测试访问url
func Test6() {
	for _, urlstr := range os.Args[1:] {
		txt, err := http.Get(urlstr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error：%v\n", err)
			//终止进程，并且返回错误代码1
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(txt.Body)
		//防止资源泄漏得到数据后需要关闭
		txt.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "URL：%s ,Error : %v\n", urlstr, err)
			//终止进程，并且返回错误代码1
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

var palette = []color.Color{color.White, color.RGBA{0x00, 0xFF, 0x00, 0xff}, color.RGBA{0xFF, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xFF, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 3 // next color in palette
)

//测试gif动画生成  $ go run test.go > test.gif
func Test5() {
	//获取19位时间戳  ,纳秒数
	//将纳秒数当作种子，种子的作用就好象一个key，使用同一个种子得到的随机数永远是一样的，计算速度根据服务器配置有关系，计算机1纳秒15次左右
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout, 5)
}

func lissajous(out io.Writer, cyc float64) {
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
		tmp := float64(cycles)
		if cyc > 0 {
			tmp = cyc
		}
		for t := 0.0; t < tmp*2*math.Pi; t += res {
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
