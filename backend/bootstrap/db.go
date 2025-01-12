package bootstrap

import (
	"github.com/lining4069/kv-auth-db/backend/app/models"
	"github.com/lining4069/kv-auth-db/backend/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

/**
gorm 有一个默认的logger，由于日志内容是输出到控制台的，
我们需要自定义一个写入器，
将默认的logger.Writer 接口的实现，切换为自定义的写入器
*/

// getGormLogWriter 创建自定义的日志写入器
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Database.EnableFileLogWriter {
		// 自定义writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 使用默认的Writer
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}

// getGormLogger 切换默认Logger,改为使用的Writer
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢SQL阈值
		LogLevel:                  logMode,                                         // 日志级别
		IgnoreRecordNotFoundError: false,                                           //忽略ErrRecordNotFound(记录未找到)错误
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, // 禁用彩色打印
	})
}

// InitializeDB 初始化数据库
func InitializeDB() *gorm.DB {
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMysqlGorm()
	default:
		return initMysqlGorm()
	}
}

// initMysqlGorm 初始化mysql+gorm
func initMysqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database

	if dbConfig.Database == "" {
		return nil
	}

	// 构建mysql连接信息
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"

	// 连接重试机制
	maxRetries := 5
	retryInterval := 10 * time.Second
	for i := 0; i < maxRetries; i++ {
		mysqlConfig := mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         191,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  //禁用datetime精度
			DontSupportRenameIndex:    true,  // 重名索引时采用删除并新建方式
			DontSupportRenameColumn:   true,  //用change 重名列
			SkipInitializeWithVersion: false, //根据版本自动配置
		}
		db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外检约束
			Logger:                                   getGormLogger(), // 使用自定义的Logger 替换默认输出到终端的logger
		})
		if err == nil {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns) // 空闲连接池中连接的最大数量
			sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns) // # 打开数据库连接的最大数量
			// 数据库初始化
			initMySqlTables(db)
			return db
		}
		global.App.Log.Error("mysql connect failed, retrying", zap.Int("attempt", i+1), zap.Error(err))
		time.Sleep(retryInterval)
	}
	global.App.Log.Error("mysql connect failed ! ")
	return nil
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}

}
