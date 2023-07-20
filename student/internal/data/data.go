package data

import (
	"github.com/hoisie/mustache"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet 第 4 步，用 wire 注入代码，修改 原来的 NewSet
var ProviderSet = wire.NewSet(NewData, NewGormDB, NewStudentRepo)

// Data 第 1 步引入 *gorm.DB
type Data struct {
	// TODO wrapped database client
	gormDB *gorm.DB
}

// NewGormDB 第 2 步初始化 gorm
func NewGormDB() (*gorm.DB, error) {
	prefix := mustache.Render("resource.database.student.wr")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetString(prefix + ".port")
	maxIdle := viper.GetInt(prefix + ".maxIdle")
	maxOpen := viper.GetInt(prefix + ".maxOpen")
	maxLifetime := viper.GetInt(prefix + ".maxLifetime")

	url := mustache.Render("{{user}}:{{password}}@tcp({{host}}:{{port}})/{{database}}?charset=utf8&parseTime=True&loc=Local", map[string]interface{}{
		"user":     user,
		"password": password,
		"database": database,
		"host":     host,
		"port":     port,
	})
	//db, err := gorm.Open(dialect, url)'
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //打印所有执行的sql语句
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true, // 使用单数表名
		//},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))
	return db, err
}

// NewData 第 3 步，初始化 Data
func NewData(logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{gormDB: db}, cleanup, nil
}
