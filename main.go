package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello World! Golang")
	// sampleReadFromString()
	analysisFile("default.conf")
}

/*
 * ReadFrom 输出
 */
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)
}

func analysisFile(filePath string) {
	// 打开文件对象
	file, err := os.Open(filePath)
	// 如果出错返回
	if nil != err {
		fmt.Println("读取文件出错...")
		return
	}
	// 运行结束关闭文件对象
	defer file.Close()
	// 获取一个文件读取对象
	fileReader := bufio.NewReader(file)
	// 行数计数
	index := 0
	// 创建待赋值的结构体
	cb := new(configBean)

	t := reflect.TypeOf(cb)
	fmt.Printf("Type:%s\n", t.Name())
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(cb)
	fmt.Printf("结构体类型是:%s\n", v)
	fmt.Println("结构体类型是: ", v)
	for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
		f := t.Field(i)                                   //取得结构体的第i个字段
		val := v.Field(i).Interface()                     //取得字段的值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //第i个字段的名称,类型,值
	}

	// 开始读取
	for {
		// 循环读取文件
		str, err := fileReader.ReadString('\n')
		index++
		strs := strFrag(str, "=")
		fmt.Printf("分割 key: %s; value: %s\n", strs[0], strs[1])
		// 读取结束循环
		if io.EOF == err {
			break
		}
	}

	fmt.Printf("获取的行数是: %d\n", index)
}

/*
 * 配置文件的字符串分割处理
 */
func strFrag(str string, mark string) []string {
	strs := strings.Split(str, mark)
	if 2 < len(strs) {
		ss := ""
		rstrs := make([]string, 2, 2)
		for i := 1; i < len(strs); i++ {
			ss += strs[i]
		}
		rstrs[0] = strs[0]
		rstrs[1] = ss
		return rstrs
	}
	return strs
}

/*
 * 将获取的结果赋值给结构体对象
 */
type configBean struct {
	name string
	size int64
}
