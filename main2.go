package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var dic map[string]interface{} = make(map[string]interface{})

func initDic() {
	f, err := os.Open("./dic.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		addWord(dic, strings.Replace(line, "\n", "", -1))
	}

}

func addWord(dic map[string]interface{}, word string) {
	tempMap := dic
	runes := []rune(word)
	for i, r := range runes {
		if len(runes) <= i+1 {
			break
		}
		char := string(r)
		nChar := string(runes[i+1])
		if tempMap[char] != nil && fmt.Sprintf("%T", tempMap[char]) != "string" {
			tempMap = tempMap[char].(map[string]interface{})
			if tempMap[nChar] == nil {
				tempMap[nChar] = "END"
			}

		} else {
			tempMap2 := make(map[string]interface{})
			tempMap2[nChar] = "END"
			tempMap[char] = tempMap2
			tempMap = tempMap2
		}
	}
}
func sensitiveFind(runes []rune) []string {
	step := 0
	var find func(dictionary map[string]interface{}, key string, result []string, level int) []string
	find = func(dictionary map[string]interface{}, key string, result []string, level int) []string {
		for step < len(runes) {
			char := string(runes[step])
			step++
			if dictionary[char] != nil {
				if fmt.Sprintf("%T", dictionary[char]) == "string" {
					result = append(result, key+char)
					return result
				}
				result = find(dictionary[char].(map[string]interface{}), key+char, result, level+1)
				if level != 0 {
					return result
				}
			} else {
				if level != 0 {
					return result
				}
			}
		}
		return result
	}
	return find(dic, "", nil, 0)
}

func main() {
	initDic() //伟大领袖毛泽东1，是我们毛ze东心中的毛泽D3红太阳,毛啊
	addWord(dic, "毛泽东")
	addWord(dic, "泽东1")
	s := "伟大领袖mao泽东1，是我们毛ze东心中的毛泽D3红太阳,毛啊潮吹"
	result := sensitiveFind([]rune(s))
	// enc := json.NewEncoder(os.Stdout)
	// enc.Encode(dic)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", sensitiveReplace([]rune(s)))
}
