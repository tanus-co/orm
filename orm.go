package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/tanus-co/snowflake/rpc/client"
)

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tanus_" + defaultTableName
	}
}

func insertId(scope *gorm.Scope) {
	if scope.HasColumn("id") {
		id := client.NewSnowflakeClient().GetId()
		if id > 0 {
			err := scope.SetColumn("id", id)
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic("Could not connect the snowflake server")
		}

	}
}

func CreateDB() *gorm.DB {
	driver := viper.GetString("orm.driver")
	dsn := viper.GetString("orm.dsn")
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(viper.GetBool("orm.debug"))
	db.Callback().Create().Before("gorm:create").Register("insert_id", insertId)
	return db
}
