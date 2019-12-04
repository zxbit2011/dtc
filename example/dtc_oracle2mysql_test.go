package example

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/oci8"
	_ "github.com/mattn/go-oci8"
	"testing"
)

type Userinfo struct {
	ID   int64  `gorm:"column:id" form:"id"`
	Name string `gorm:"column:name" form:"name"`
}

var driverName = "oci8"
var dataSourceName = "system/123456@127.0.0.1:1521/ORCL"

func TestOracle2MySQL(t *testing.T) {
	db, err := gorm.Open(driverName, dataSourceName)
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
}
