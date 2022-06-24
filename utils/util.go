package utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// 去除空字符
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}

// 文件创建
// 创建测试数据
func CreateFile(text *string, fileName string, fileType string) {

	create, err := os.Create("C:\\Users\\Administrator\\Desktop\\" + fileName + fileType)

	if err != nil {
		fmt.Println(err)
	}

	defer create.Close()

	writeString, err := create.WriteString("[" + *text + "]")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("字节:" + strconv.Itoa(writeString))

}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ConvertStr2GBK(str string) string {
	//将utf-8编码的字符串转换为GBK编码
	ret, _ := simplifiedchinese.GBK.NewEncoder().String(str)
	return ret //如果转换失败返回空字符串

	//如果是[]byte格式的字符串，可以使用Bytes方法
	b, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	return string(b)
}

func ConvertGBK2Str(gbkStr string) string {
	//将GBK编码的字符串转换为utf-8编码
	ret, _ := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
	return ret //如果转换失败返回空字符串

	//如果是[]byte格式的字符串，可以使用Bytes方法
	b, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(gbkStr))
	return string(b)
}
