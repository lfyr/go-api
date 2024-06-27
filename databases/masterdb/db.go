package masterdb

import (
	"github.com/lfyr/go-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	m := utils.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		logrus.Error(m)
		return
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                            // 慢 SQL 阈值
			LogLevel:                  logger.LogLevel(m.GetLevel(m.LogMode)), // 日志级别
			IgnoreRecordNotFoundError: false,                                  // 忽略记录未找到错误
			Colorful:                  true,                                   // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Error(m.Dsn(), err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	}
	DB = db
}
