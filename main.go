package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"httpParse/db"
	"httpParse/guojiayibao"
	"httpParse/hs"
	"httpParse/protos"
	"sync"
	"time"
)

// protobuf 测试
func protobufTest() {
	s1 := &protos.Student{} // 第一个测试
	s1.Name = "jz01"
	s1.Age = 23
	s1.Address = "cq"
	s1.Cn = protos.ClassName_class2 //枚举类型赋值
	ss := &protos.Students{}
	ss.Person = append(ss.Person, s1) //将第一个学生信息添加到Students对应的切片中
	s2 := &protos.Student{}           //第二个学生信息
	s2.Name = "jz02"
	s2.Age = 25
	s2.Address = "cd"
	s2.Cn = protos.ClassName_class3
	ss.Person = append(ss.Person, s2) //将第二个学生信息添加到Students对应的切片中
	ss.School = "cqu"
	fmt.Println("Students信息为：", ss)
	// Marshal takes a protocol buffer message
	// and encodes it into the wire format, returning the data.
	buffer, _ := proto.Marshal(ss)
	fmt.Println("序列化之后的信息为：", buffer)
	// Use UnmarshalMerge to preserve and append to existing data.
	data := &protos.Students{}
	proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)
}

func T1001() {
	for i := 1; i < 22; i++ {
		data := guojiayibao.GetData1001(i)
		configure, _ := db.MysqlConfigure()
		for _, row := range data.Rows {
			if row.AuditAddTime.IsZero() {
				row.AuditAddTime = time.Now()
			}
			configure.Table("t_drug1001_tmp").Create(&row)
		}
	}
}

func T1002() {
	for i := 1; i < 44; i++ {
		data1002 := guojiayibao.GetData1002(i)
		configure, _ := db.MysqlConfigure()
		for _, data := range data1002.Rows {
			configure.Table("t_drug_info1002_tmp").Create(&data)
		}
	}
}

func T1003() {
	for i := 1; i < 30; i++ {
		data := guojiayibao.GetData1003(i)
		configure, _ := db.MysqlConfigure()
		for _, row := range data.Rows {
			configure.Table("t_good_info1003").Create(&row)
		}
	}
}

func Hs(tag int) {
	for i := 2; i < 30; i++ {
		hs.ExampleScrape(tag, i)
	}
}

// 全局变量
var wg sync.WaitGroup

// 5203
var tag = 1

func flush() {
	defer wg.Done()
	for i := 213; i < 1917; i++ {
		hs.Paoyou(tag, i)
	}
}

func Tpaoyou(num1, num2 int) {
	defer wg.Done()
	for i := num1; i < num2; i++ {
		hs.Paoyou(tag, i)
	}
}

func newPaoYou(num1, num2, size int) {
	if num2-num1 > 0 {
		wg.Add(1)
		for i := 1; i <= size; i++ {
			go Tpaoyou(num1+(num2-num1)/size*(i-1), num1+(num2-num1)/size*i)
		}
		wg.Wait()
	} else {
		fmt.Printf("num2 - num1 > 0 , 修改参数")
	}
}

// 1000 - 2000 1000-1333 1333-1666 1666-2000
func main() {
	//gin := gin.Default()
	//gin.GET("/getData/:page/:start", src.SaveInfo)
	//gin.Run(":8500")
	//db.AutoCreateTable(xml.XmlInfo{})
	//redisSetGet()
	//for i := 2; i <= 8; i++ {
	//	hs.RedisLi5apuu7(22, i)
	//}
	hs.Redis2Mysql()
}
