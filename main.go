package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World! Golang...")
	cb, err := analysisFile("default.conf")
	if nil != err {
		fmt.Printf("出错了...=>%v", err)
		return
	}
	fmt.Printf("解析配置文件:第一个值[%v],第二个值[%v].\n", cb.Name, cb.Size)
}

func analysisFile(filePath string) (ConfigBean, error) {
	// 创建待赋值的结构体
	cb := ConfigBean{}
	// 打开文件对象
	file, err := os.Open(filePath)
	// 如果出错返回
	if nil != err {
		fmt.Println("读取文件出错...")
		return cb, err
	}
	// 运行结束关闭文件对象
	defer file.Close()
	// 获取一个文件读取对象
	fileReader := bufio.NewReader(file)
	// 行数计数
	index := 0

	t := reflect.TypeOf(cb)

	c := reflect.ValueOf(&cb).Elem()

	kv := make([][]string, 0, 10)

	// 开始读取
	for {
		// 循环读取文件
		str, err := fileReader.ReadString('\n')
		index++
		strs := strFrag(str, "=")
		kv = append(kv, strs)
		// 读取结束循环
		if io.EOF == err {
			break
		}
	}
	//NumField取出这个接口所有的字段数量
	for i := 0; i < t.NumField(); i++ {
		//取得结构体的第i个字段
		f := t.Field(i)
		switch f.Type.Kind() {
		case reflect.String:
			for _, v := range kv {
				if strings.EqualFold(f.Name, v[0]) {
					c.FieldByName(v[0]).SetString(v[1])
				}
			}
		case reflect.Int64:
			for _, v := range kv {
				if strings.EqualFold(f.Name, v[0]) {
					i, _ := strconv.ParseInt(v[1], 10, 64)
					c.FieldByName(v[0]).SetInt(i)
				}
			}
		case reflect.Bool:
			fmt.Printf("布尔类型\n")
		}
	}
	return cb, nil
}

/*
 * 配置文件的字符串分割处理
 */
func strFrag(str string, mark string) []string {
	strs := strings.Split(str, mark)
	strs[1] = strings.Trim(strs[1], "\n")
	if 2 < len(strs) {
		ss := ""
		rstrs := make([]string, 2, 2)
		for i := 1; i < len(strs); i++ {
			ss += strs[i]
		}
		rstrs[0] = strs[0]
		rstrs[1] = strings.Trim(ss, "\n")
		return rstrs
	}
	return strs
}

/*
 * 将获取的结果赋值给结构体对象
 */
type ConfigBean struct {
	Name string
	Size int64
}
