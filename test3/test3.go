package test3

import (
	"fmt"
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

//测试uint8结果溢出最大值，溢出后重新计算值
func Test3() {
	var u uint8 = 255
	fmt.Println(u, u+1, u+2, u+10)
}

//测试数组下标溢出赋值
func Test4() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]
	// var s [3]string
	// var ss []string = s[1:]
	// for i := 0; i < 3; i++ {
	// 	//s[i] = "sss" + strconv.Itoa(i)
	// 	append(ss, "sss")
	// }
	mySlice = append(mySlice, 1)
	//mySlice[9] = 100  //Error
	mySlice[0] = 100 //Success
	//mySlice = append(mySlice, 1)
	//mySlice = append(mySlice, 1)
	fmt.Println(mySlice, len(mySlice), cap(mySlice))
}

//测试通道  后进先出
func Test5() {
	ch := make(chan int, 2)
	go sum(5, ch)
	go sum(10, ch)
	go sum(100, ch)

	fmt.Println("1", <-ch)
	fmt.Println("2", <-ch)
	fmt.Println("3", <-ch)
}

func sum(x int, ch chan int) {
	s := 0
	for i := 1; i <= x; i++ {
		s += i
	}
	ch <- s
}

func Test6() {
	strs := []string{"aaa", "bbb", "ccc"}
	for i := len(strs) - 1; i >= 0; i-- {
		fmt.Println(strs[i])
	}
}
