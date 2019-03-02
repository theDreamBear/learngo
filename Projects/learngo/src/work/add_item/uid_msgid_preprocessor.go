package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func WriteFile(result map[string]string) error {
	gFiles := make([]io.Writer, 100)

	for key, value := range result {
		if intKey, err := strconv.ParseInt(key, 10, 64); err == nil {
			idx := intKey % 100
			file := filename + strconv.FormatInt(idx, 10)
			line := key + "\t" + value + "\n"
			// 文件存在
			if gFiles[idx] != nil {
				gFiles[idx].Write([]byte(line))
			} else {

				if fd, err := os.Create(ParentPath + file); err != nil {
					panic(err)
				} else {
					gFiles[idx] = fd
					fd.Write([]byte(line))
					defer fd.Close()
				}
			}
		} else {
			panic(err)
		}
	}
	return nil
}

func StringSplit(str string) []string {
	arr := strings.Split(str, ",")
	var (
		result []string
	)
	for _, v := range arr {
		if v != "" {
			v = strings.Trim(v, " ")
			result = append(result, v)
		}
	}
	return result
}

func LineSplit(line string) (string, string, []string) {
	arr := strings.Split(line, "\t")
	addtime := arr[0]
	msgid := arr[1]
	uids := StringSplit(arr[2])
	return addtime, msgid, uids
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadLine(path string) {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	result := make(map[string]string)
	number := 0
	for {
		number = number + 1
		if number%1000000 == 0 {
			fmt.Println(number)
		}
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		addtime, msgid, uids := LineSplit(string(a))
		//fmt.Println(msgid)
		for _, uid := range uids {
			if v, ok := result[uid]; ok {
				//strings.Join([]string{v, msgid}, "\t")
				result[uid] = v + "\t" + msgid + "_" + addtime
				//fmt.Println(v)
			} else {
				result[uid] = msgid + "_" + addtime
			}
		}
	}
	WriteFile(result)
	//for k, v := range result {
	//	fmt.Println(k, v)
	//}
}

func FileRead(path string) (int, error) {
	if ok, err := PathExists(path); err != nil {
		//fmt.Println(err)
		return -1, errors.New("未知错误")
	} else {
		if ok == false {
			//	fmt.Println("file no exist")
			return -1, errors.New("文件不存在")
		}
		ReadLine(path)
	}

	return 0, nil
}

func pathParse(path string) {
	idx := strings.LastIndex(path, "/")
	ParentPath = path[:idx]
	ParentPath += tailPath
	//fmt.Println(ParentPath)
	exist, err := PathExists(ParentPath)
	if err != nil {
		panic(err)
	} else {
		if exist == false {
			err := os.MkdirAll(ParentPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

var tailPath string = "/processed_data/add_item/"
var ParentPath string
var filename string = "add_"

func main() {
	args := os.Args
	path := args[1]
	//path := "/Users/nichao/Code/GoWork/Projects/learngo/src/work/test"
	pathParse(path)
	FileRead(path)
	//fmt.Printf("%v", "9223372036854775807")
	//result := make(map[string]string)
	//fmt.Println(result)
}
