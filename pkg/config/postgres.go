package config

import (
	"fmt"
	"github.com/google/martian/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	PgDB *gorm.DB
)

func InitPostgres(database Database) error {
	var err error
	dbCfg := database
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s TimeZone=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Database,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.SslMode,
		dbCfg.Timezone,
	)
	dbConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), // 自定义日志打印
		DisableForeignKeyConstraintWhenMigrating: true,                                // 关闭外键强关联
		PrepareStmt:                              false,                               // 禁用隐式prepared statement预编译
		DisableAutomaticPing:                     false,                               // 禁用自动ping可用性
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}
	PgDB, err = gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	log.Infof("%v", dsn+"连接成功！")

	PgDB.Logger = logger.Default.LogMode(logger.Info)
	sqlDB, err := PgDB.DB()
	if err != nil {
		log.Errorf("获取sqlDB失败: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(dbCfg.MaxIdle)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpen)
	sqlDB.SetConnMaxIdleTime(dbCfg.MaxIdleTime)
	sqlDB.SetConnMaxLifetime(dbCfg.MaxLifeTime)

	err = PgDB.AutoMigrate( /*&model.CdkmInspectionRecord{}*/ )
	return err
}
