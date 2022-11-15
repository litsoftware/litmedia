package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/litsoftware/litmedia/internal/pkg/config"
)

func init() {
	_ = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		config.GetString("mysql.default.user"),
		config.GetString("mysql.default.password"),
		config.GetString("mysql.default.host"),
		config.GetString("mysql.default.port"),
		config.GetString("mysql.default.database"))
}
