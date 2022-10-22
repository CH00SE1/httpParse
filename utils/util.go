package utils

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

// 判断文件是否存在 不存在创建
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		return true, nil
	}
	return false, err
}

// 文件创建
func CreateFile(text *string, path, fileName, fileType string) {

	// 判断文件是否创建
	pathExists(path)

	create, err := os.Create(path + fileName + fileType)

	if err != nil {
		fmt.Println(err)
	}

	defer create.Close()

	writeString, err := create.WriteString(*text)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("字节:%d - 文件名称:%s\n", writeString, fileName)

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

// 图片保存
func SavePhoto(url, fileName string, i int) {
	log.Println("url : " + url)
	response, _ := http.Get(url)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	basePath := "C:\\Users\\Administrator\\Desktop\\photo_lilin\\"
	pathExists(basePath)
	name := basePath + strconv.Itoa(i+1) + "[学校照片]" + UrlDecode(fileName)
	if !strings.Contains(name, ".jpg") {
		name += ".jpg"
	}
	log.Println("path : " + name)
	WriteFile(name, bytes, 0755)
}

// 写入文件
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// 解密base64str
func Base64ToStr(str string) {
	var text *string
	decodeString, err := base64.StdEncoding.DecodeString(str)
	bytes2String := Bytes2String(decodeString)
	text = &bytes2String
	if err != nil {
		fmt.Println("Error", err)
	}
	CreateFile(text, "", "login", ".txt")
}

// Unicode加密
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// Unicode解密
func UrlDecode(str string) string {
	res, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return res
}

// 拿到正式m3u8视频下载地址
//func IsM3u8Ture(url string) string {
//	response, _ := http.Get(url)
//	buf := new(bytes.Buffer)
//	buf.ReadFrom(response.Body)
//	newStr := buf.String()
//	if buf == nil {
//		return url
//	}
//	if strings.Contains(newStr, ".ts") {
//		return url
//	}
//	index := strings.Index(url, "/20")
//	url1 := url[:index]
//	if strings.Contains(newStr, "/") {
//		index2 := strings.Index(newStr, "/")
//		openUrl := url1 + newStr[index2:]
//		return openUrl
//	}
//	return url
//}
