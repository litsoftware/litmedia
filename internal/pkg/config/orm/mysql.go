package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	_ "gorm.io/driver/mysql"
	"log"
)

var err error

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.GetString("mysql.default.user"),
		config.GetString("mysql.default.password"),
		config.GetString("mysql.default.host"),
		config.GetString("mysql.default.port"),
		config.GetString("mysql.default.database"))

	client, err = ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %s %v", dsn, err)
	}
}
