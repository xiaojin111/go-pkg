package mysqldb

import (
	"fmt"
	"time"

	// import mysql driver fo gorm
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinmukeji/go-pkg/log/gormlogger"
	"github.com/jinzhu/gorm"
)

// DbClient 是数据访问管理器
type DbClient struct {
	*gorm.DB
	opts Options
}

func Open(options Options) (*gorm.DB, error) {
	// mysql 连接字符串格式:
	// 	`username:password@tcp(localhost:3306)/db_name?charset=utf8mb4&parseTime=True&loc=utc`
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		options.Username,
		options.Password,
		options.Address,
		options.Database,
		options.Charset,
		options.ParseTime,
		options.Locale)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// gorm setting
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(options.MaxConnections)
	db.SetLogger(gormlogger.New(options.Address, options.Database))
	db = db.LogMode(options.EnableLog)
	// 禁止没有 WHERE 语句的 DELETE 或 UPDATE 操作执行，否则抛出 error
	db = db.BlockGlobalUpdate(options.BlockGlobalUpdate)
	// 重置 SetNow 的时间获取方式为总是获取UTC时区时间
	db = db.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})

	return db, nil
}

// NewDbClient 根据传入的 options 返回一个新的 DbClient
func NewDbClient(opts ...Option) (*DbClient, error) {
	options := NewOptions(opts...)
	db, err := Open(options)
	if err != nil {
		return nil, err
	}

	return &DbClient{db, options}, nil
}

// Options 返回 DbClient 的 Options.
func (c *DbClient) Options() Options {
	return c.opts
}
