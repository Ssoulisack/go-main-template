package bootstrap

import (
	"fmt"
	"kkl-v2/core/logs"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// **************************************** Mysql ****************************************************
func NewDatabaseConnection(env *Env) *gorm.DB {
	dsnMaster := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.Database.MasterUsername,
		env.Database.MasterPassword,
		env.Database.MasterHost,
		env.Database.MasterPort,
		env.Database.MasterDBName,
	)

	masterDB, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect master database: %v", err)
	}
	logs.Info("database connection success master")

	// Master pool settings
	sqlDB, _ := masterDB.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Only configure replica if replica config is provided
	if env.Database.Status {
		dsnReplica := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.Database.ReplicaUsername,
			env.Database.ReplicaPassword,
			env.Database.ReplicaHost,
			env.Database.ReplicaPort,
			env.Database.ReplicaDBName,
		)

		err = masterDB.Use(
			dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{
					mysql.Open(dsnReplica),
				},
				Policy: dbresolver.RandomPolicy{},
			}).
				SetMaxIdleConns(10).
				SetMaxOpenConns(50).
				SetConnMaxLifetime(time.Hour),
		)
		if err != nil {
			log.Fatalf("failed to register replica: %v", err)
		}
		logs.Info("replica database configured successfully")
	} else {
		logs.Info("no replica configuration found, using master for all operations")
	}

	return masterDB
}

func Migrate(db *gorm.DB) {
	err := db.Clauses(dbresolver.Write).AutoMigrate(
	// Your entities
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	logs.Info("Migrate successfully")
}
