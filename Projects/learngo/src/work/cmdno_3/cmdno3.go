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

func WriteFile(uid string, result string) error {
	if intKey, err := strconv.ParseInt(uid, 10, 64); err == nil {
		idx := intKey % 100
		if gFiles[idx] != nil {
			gFiles[idx].Write([]byte(result))
		} else {
			fmt.Println("io.reader == nil")
		}
	} else {
		panic(err)
	}
	return nil
}

func msgidSplit(str string) string {
	arr := strings.Split(str, ":")
	return arr[0]
}

func StringSplit(str string) []string {
	arr := strings.Split(str, ",")
	var (
		result []string
	)
	for _, v := range arr {
		if v != "" {
			v = strings.Trim(v, " ")
			//if uid, err := strconv.ParseInt(v, 10, 64); err == nil {
			//	result = append(result, uid)
			//} else {
			//	panic("cannot convert to int64")
			//}
			msgid := msgidSplit(v)
			result = append(result, msgid)
		}
	}
	return result
}

func LineSplit(line string) (string, []string) {
	arr := strings.Split(line, "\t")
	//msgid, err := strconv.ParseInt(arr[0], 10, 64)
	//if err != nil {
	//	panic("cannot convert to int64")
	//}
	//uids := StringSplit(arr[1])
	uid := arr[0]
	msgids := StringSplit(arr[1])
	return uid, msgids
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
	result := ""
	uid := ""

	for idx := int64(0); idx < 100; idx++ {
		file := filename + strconv.FormatInt(idx, 10)
		if fd, err := os.Create(ParentPath + file); err != nil {
			panic(err)
		} else {
			gFiles[idx] = fd
			defer fd.Close()
		}
	}

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		uid, msgids := LineSplit(string(a))

		result = uid + "\t" + strings.Join(msgids, "\t") + "\n"

	}
	WriteFile(uid, result)
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

var gFiles = make([]io.Writer, 100)
var tailPath = "/processed_data/cmdno_3/"
var ParentPath string
var filename = "pull_"

func main() {
	args := os.Args
	path := args[1]
	pathParse(path)
	FileRead(path)
	//
	//fmt.Printf("%v", []byte("9223372036854775807"))
	//result := make(map[string]string)
	//fmt.Println(result)
}
