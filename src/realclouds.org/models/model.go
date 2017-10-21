package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pborman/uuid"
)

//Model *
type Model struct {
	ID        string     `sql:"index" gorm:"primary_key;column:id;type:varchar(100)" json:"id" xml:"id"`
	CreatedAt time.Time  `sql:"index" gorm:"column:created_at;type:timestamp" json:"created_at" xml:"created_at"`
	UpdatedAt time.Time  `sql:"index" gorm:"column:updated_at;type:timestamp" json:"updated_at" xml:"updated_at"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at;type:timestamp" json:"deleted_at" xml:"deleted_at"`
}

//BeforeCreate ID处理
func (d *Model) BeforeCreate(scope *gorm.Scope) error {
	uuidStr := uuid.NewRandom().String()
	if err := scope.SetColumn("ID", uuidStr); nil != err {
		return err
	}
	return nil
}

//DBCtx *
type DBCtx struct {
	WebContext echo.Context
}

//NewDBCtx *
func NewDBCtx(c echo.Context) *DBCtx {
	dbCtx := &DBCtx{WebContext: c}
	return dbCtx
}

//MySQL *
func (d *DBCtx) MySQL() *gorm.DB {
	return d.WebContext.Get("mysql").(*gorm.DB)
}

//AutoMigrate * AutoMigrate
func AutoMigrate(c echo.Context, schema ...interface{}) error {
	return NewDBCtx(c).MySQL().AutoMigrate(schema...).Error
}
