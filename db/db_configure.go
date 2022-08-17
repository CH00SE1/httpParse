package db

import (
	"database/sql"
	"fmt"
	"github.com/qmhball/db2gorm/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *sql.DB
var ERP, DS = "erp", "ds"

// 连接配置
type connect struct {
	HOST      string
	PORT      string
	DATABASE  string
	USERNAME  string
	PASSWORD  string
	CHARSET   string
	PARSETIME string
	SID       string
}

var erp = connect{
	HOST:     "192.168.191.212",
	PORT:     "1521",
	USERNAME: "dsjkldjl",
	PASSWORD: "dhsjkahkdj",
	SID:      "data",
}

var ds = connect{
	HOST:     "192.168.191.41",
	PORT:     "1521",
	USERNAME: "djskjdl",
	PASSWORD: "akjshdklhj",
	SID:      "djskjdlks",
}

/**
 * @title 数据库配置
 * @author xiongshao
 * @date 2022-06-24 10:15:21
 */
// 数据库1号
const dsn1 = "root:xiAtiAn@djwk@tcp(192.168.10.142:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local"

// localhost
const dsn3 = "root:11098319@tcp(192.168.10.87:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local"

// gongsi
const dsn = "root:11098319@tcp(192.168.10.87:3306)/local?charset=utf8mb4&parseTime=True&loc=Local"

// MySQL驱动高级配置
func MysqlConfigure() (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		// string 类型字符默认长度
		DefaultStringSize: 256,
		// 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DisableDatetimePrecision: true,
		// 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameIndex: true,
		// 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		DontSupportRenameColumn: true,
		// 根据当前 MySQL 版本自动配置
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		// 跳过默认事物
		SkipDefaultTransaction: true,
		// GORM 允许用户通过覆盖默认的NamingStrategy来更改命名约定
		NamingStrategy: schema.NamingStrategy{
			// 表名前缀，`User`表为`t_users`
			TablePrefix: "t_",
			// 使用单数表名，启用该选项后，`User` 表将是`user`
			SingularTable: true,
		},
		// 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: true,
		// 更改创建时间使用的函数
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// url : https://www.zhihu.com/question/490032130/answer/2299150794
func AutoCreateTable(Table interface{}) {
	db, _ := MysqlConfigure()
	// 创建表
	//db.Migrator().CreateTable(&Table)
	// 将 "ENGINE=InnoDB" 添加到创建 的 SQL 里去
	if !db.Migrator().HasTable(&Table) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb3").AutoMigrate(&Table)
		fmt.Println("创建成功")
	} else {
		fmt.Println("表已存在")
	}
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

// erp连接
func ErpOracleConfig() (baseDB *sql.DB, err error) {
	db, err := sql.Open("godror", `user="`+erp.USERNAME+`" password="`+erp.PASSWORD+`" connectString="`+erp.HOST+`:`+erp.PORT+`/`+erp.SID+`"`)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db, nil
}

// ds连接
func DsOracleConfig() (baseDB *sql.DB, err error) {
	db, err := sql.Open("godror", `user="`+ds.USERNAME+`" password="`+ds.PASSWORD+`" connectString="`+ds.HOST+`:`+ds.PORT+`/`+ds.SID+`"`)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db, nil
}
