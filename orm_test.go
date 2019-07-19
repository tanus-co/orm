package orm

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/tanus-co/common/config"
	"github.com/tanus-co/snowflake/rpc/client"
	"testing"
)

type User struct {
	TenantModel
	Name  string `json:",omitempty"`
	IsMan bool   `json:",omitempty"`
	Age   string `json:",omitempty"`
	Age3  int
}

func TestOrm_Query(t *testing.T) {
	db := CreateDB()
	if db == nil {
		t.Error("db is nil")
	}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&User{})

	user := &User{
		TenantModel: TenantModel{
			Tenant: client.NewSnowflakeClient().GetId(),
		},
		Name:  "username",
		IsMan: false,
		Age3:  10,
	}

	db.Create(user)

	u := User{}
	db.Where("id = ?", user.ID).Find(&u)
	t.Log(u.DeletedAt)
	t.Log(u.UpdatedAt)
	j, _ := json.Marshal(u)
	t.Log(fmt.Sprintf("%s", j))

}
