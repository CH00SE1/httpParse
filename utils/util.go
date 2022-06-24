package utils

import (
	"fmt"
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
