package middleware

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Gorm 支持
	"github.com/labstack/echo"
)

//DBConfig *
type DBConfig struct {
	// Addr *
	Addr string `json:"addr" xml:"addr"`

	// MaxIdleConns *
	MaxIdleConns int `json:"max_idle_conns" xml:"max_idle_conns"`

	// MaxOpenConns *
	MaxOpenConns int `json:"max_open_conns" xml:"max_open_conns"`

	// Username *
	UserName string `json:"username" xml:"username"`

	// Password *
	Password string `json:"password" xml:"password"`

	// DataBase *
	DataBase string `json:"database" xml:"database"`

	//LogMode *
	LogMode bool `json:"log_mode" xml:"log_mode"`
}

//DefaultConfig *
var DefaultConfig = DBConfig{
	Addr: "127.0.0.1:3306",

	// MaxIdleConns *
	MaxIdleConns: 10,
	// MaxOpenConns *
	MaxOpenConns: 100,

	// Username *
	UserName: "root",

	// Password *
	Password: "123456",

	// DataBase *
	DataBase: "test",

	//LogMode *
	LogMode: false,
}

//MySQLConf MySQL config
func MySQLConf(config DBConfig) (*gorm.DB, error) {

	if len(config.Addr) == 0 {
		config.Addr = DefaultConfig.Addr
	}

	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = DefaultConfig.MaxIdleConns
	}

	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = DefaultConfig.MaxOpenConns
	}

	if len(config.UserName) == 0 {
		config.UserName = DefaultConfig.UserName
	}

	if len(config.Password) == 0 {
		config.Password = DefaultConfig.Password
	}

	if len(config.DataBase) == 0 {
		config.DataBase = DefaultConfig.DataBase
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Asia%%2FShanghai&timeout=30s", config.UserName, config.Password, config.Addr, config.DataBase)

	db, err := gorm.Open("mysql", dbURL)
	if nil != err {
		return nil, err
	}

	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	db.LogMode(config.LogMode)

	if err = db.DB().Ping(); nil != err {
		return nil, err
	}

	return db, nil
}

//DB *
type DB struct {
	Gorm  *gorm.DB
	Mutex sync.RWMutex
}

//MwMySQL MySQL middleware
func MwMySQL(db *gorm.DB) *DB {
	return &DB{Gorm: db}
}

//MySQL Server header
func (d *DB) MySQL(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		d.Mutex.Lock()
		defer d.Mutex.Unlock()
		c.Set("mysql", d.Gorm)
		return next(c)
	}
}
