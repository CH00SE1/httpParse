package utils

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/qmhball/db2gorm/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 去除空字符
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}

// 数据库url
const dsn = "root:xiAtiAn@djwk@tcp(192.168.10.142:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local"

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

// mysql数据库连接访问
// 夏添 - 主机
func DB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印执行日志
		Logger: logger.New(
			log.New(colorable.NewColorableStdout(), "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 开启彩色打印
			}),
	})
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

// 结构体生成器
func MysqlGen(TableName string) {
	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, TableName)
}
